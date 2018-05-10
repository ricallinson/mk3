package main

import (
	"flag"
	"fmt"
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
	flag.StringVar(&dongle, "dongle", "", "Serial port that's connected to the Dongle Terminator.")
	var raw string
	flag.StringVar(&raw, "raw", "", "Send a raw command to the Dongle Terminator.")
	var scanCells bool
	flag.BoolVar(&scanCells, "scan-cells", false, "Scans the bus for cells.")
	var scanCards bool
	flag.BoolVar(&scanCards, "scan-cards", false, "Scans the bus for cards.")
	var setup bool
	flag.BoolVar(&setup, "setup", false, "Walks through assigning addresses to all cards in the BMS.")
	var commands string
	flag.StringVar(&commands, "cmd", "", "Path to the YAML configuration file of commands to execute.")
	var addr int
	flag.IntVar(&addr, "addr", -1, "The address to which the commands are to be executed. Default is all.")
	var maxAddr int
	flag.IntVar(&maxAddr, "max-addr", 255, "The highest address to which the commands are to be executed. Default is 255.")
	var newAddr int
	flag.IntVar(&newAddr, "new-addr", 0, "Changes the address of the ONLY card on attached to the Dongle.")
	var volts bool
	flag.BoolVar(&volts, "volts", false, "Scans the bus and returns the average voltage for all cells found.")
	var temps bool
	flag.BoolVar(&temps, "temps", false, "Scans the bus and returns the average temperature for all cells found.")
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
		os.Exit(scanForCells(mk3DT))
	}
	if scanCards {
		os.Exit(scanForCards(mk3DT))
	}
	if setup {
		os.Exit(setupBus(mk3DT))
	}
	if newAddr > 0 {
		os.Exit(setAddr(mk3DT, newAddr))
		return
	}
	if volts {
		os.Exit(listVolts(mk3DT))
		return
	}
	if temps {
		os.Exit(listTemps(mk3DT))
		return
	}
	if commands == "" {
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

// Prints to standard out the result.
func sendRawCommand(mk3DT *Mk3DT, s string) {
	log.Println(string(mk3DT.Raw(s + "\n\r")))
}

// Prints to standard out the result.
func scanForCells(mk3DT *Mk3DT) int {
	pos := 1
	lastSn := 0
	lastEmpty := false
	for addr := 1; addr <= 255; addr++ {
		if sn := mk3DT.GetSerialNum(addr); sn > 0 {
			if sn == lastSn {
				pos++
			} else {
				pos = 1
			}
			if lastEmpty {
				fmt.Println("")
			}
			fmt.Printf("Cell %03d on card %05d at position %03d\n\r", addr, sn, pos)
			lastEmpty = false
			lastSn = sn
		} else {
			lastEmpty = true
			fmt.Print(".")
		}
	}
	fmt.Println("Done")
	return 0
}

func scanForCards(mk3DT *Mk3DT) int {
	lastSn := 0
	count := 0
	for addr := 1; addr <= 255; addr++ {
		sn := mk3DT.GetSerialNum(addr)
		if sn != lastSn && sn != 0 {
			fmt.Printf("Starting cell %03d on %05d with %d cells\n\r", addr, sn, mk3DT.GetNumCells(addr))
			lastSn = sn
			count++
		}
	}
	fmt.Println(count, "BMS cards found")
	return 0
}

func setAddr(mk3DT *Mk3DT, newAddr int) int {
	for addr := 1; addr <= 255; addr++ {
		fmt.Print(".")
		if sn := mk3DT.GetSerialNum(addr); sn > 0 {
			mk3DT.ChangeAddr(addr, newAddr)
			fmt.Printf("\n\rAddress set to %03d\n\r", newAddr)
			return 0
		}
	}
	fmt.Println("")
	return 1
}

func listVolts(mk3DT *Mk3DT) int {
	for addr := 1; addr <= 255; addr++ {
		if v := mk3DT.GetRealTimeVoltage(addr); v > 0 {
			sn := mk3DT.GetSerialNum(addr)
			fmt.Printf("Cell %03d at %.2f VDC on card %05d\n\r", addr, v/float32(mk3DT.GetNumCells(addr)), sn)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("")
	return 0
}

func listTemps(mk3DT *Mk3DT) int {
	for addr := 1; addr <= 255; addr++ {
		if t := mk3DT.GetCellsTemp(addr); t > 0 {
			sn := mk3DT.GetSerialNum(addr)
			fmt.Printf("Cell %03d at %dc on card %05d\n\r", addr, t, sn)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("")
	return 0
}
