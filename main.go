package main

import (
	"flag"
	"github.com/tarm/serial"
	"log"
	"os"
	"time"
)

type SerialPort interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Flush() error
	Close() error
}

func main() {
	var dongle string
	flag.StringVar(&dongle, "dongle", "", "Serial port that's connected to the Dongle Terminator (required).")
	var raw string
	flag.StringVar(&raw, "raw", "", "Send a raw command to the Dongle Terminator.")
	var scanCells bool
	flag.BoolVar(&scanCells, "scan-cells", false, "Scans the bus for cells.")
	var scanCards bool
	flag.BoolVar(&scanCards, "scan-cards", false, "Scans the bus for cards.")
	var commands string
	flag.StringVar(&commands, "cmd", "", "Path to the YAML configuration file of commands to execute.")
	var addr int
	flag.IntVar(&addr, "addr", -1, "The address to which the commands are to be executed. Default is all.")
	var minAddr int
	flag.IntVar(&minAddr, "min-addr", 0, "The lowest address to which the commands are to be executed. Default is 0.")
	var maxAddr int
	flag.IntVar(&maxAddr, "max-addr", 255, "The highest address to which the commands are to be executed. Default is 255.")
	var newAddr int
	flag.IntVar(&newAddr, "new-addr", 0, "Changes the address of the first card found attached to the Dongle.")
	var realtime bool
	flag.BoolVar(&realtime, "realtime", false, "Constantly scans the bus and returns current volts, current temperature, card serial number and Number of cells on card.")
	flag.Parse()

	// Create an instance of a Mk3 Dongle Terminator.
	mk3DT := newMk3DT(dongle)
	defer mk3DT.Close()

	// Process CLI Options.
	if raw != "" {
		sendRawCommand(mk3DT, raw)
		return
	}
	if scanCells {
		os.Exit(scanForCells(mk3DT, minAddr, maxAddr))
	}
	if scanCards {
		os.Exit(scanForCards(mk3DT, maxAddr))
	}
	if newAddr > 0 {
		os.Exit(setAddr(mk3DT, newAddr))
		return
	}
	if realtime {
		mk3DT.RealtimeValues(maxAddr)
		return
	}
	if commands == "" {
		// If we get to here and there's no commands we can't do anything.
		log.Println("You must provide a path to YAML file with the commands to execute")
		return
	}
	if addr < -1 || addr > 255 {
		log.Println("The address must be between 0 and 255.")
		return
	}

	// Execute the loaded command object.
	e := NewExecutor(mk3DT)
	e.Commands = readYamlFileToExecutorCommands(commands)
	var yaml []byte
	if addr >= 0 && addr <= 255 {
		yaml = interfaceToYaml(e.ExecuteCommandsAtAddr(addr))
	} else {
		yaml = interfaceToYaml(e.ExecuteCommands(maxAddr))
	}
	// Prints to standard out the result.
	log.Println(string(yaml))
}

// Opens the serial port.
func connectToDongle(path string) (error, *serial.Port) {
	c := &serial.Config{
		Name:        path,
		Baud:        9600,
		Size:        8,
		Parity:      serial.ParityNone,
		StopBits:    serial.Stop1,
		ReadTimeout: time.Millisecond,
	}
	s, err := serial.OpenPort(c)
	if err != nil {
		return err, nil
	}
	return nil, s
}

func newMk3DT(dongle string) *Mk3DT {
	// Open Serial Port
	var serialPort SerialPort
	var serialError error
	if dongle == "" {
		serialPort = NewMockPort()
	} else {
		serialError, serialPort = connectToDongle(dongle)
	}
	if serialError != nil {
		log.Print("Error opening port: ")
		log.Fatal(serialError)
	}
	return NewMk3DT(serialPort)
}
