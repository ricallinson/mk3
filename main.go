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
	var maxAddr int
	flag.IntVar(&maxAddr, "max-addr", 255, "The highest address to which the commands are to be executed. Default is 255.")
	var newAddr int
	flag.IntVar(&newAddr, "new-addr", 0, "Changes the address of the first card found attached to the Dongle.")
	var realtime bool
	flag.BoolVar(&realtime, "realtime", false, "Constantly scans the bus and returns current volts, current temperature, card serial number and Number of cells on card.")
	flag.Parse()

	// We can't do anything without a serial port.
	if dongle == "" {
		log.Println("You must provide the serial port that's connected to the Dongle Terminator")
		return
	}

	// Create an instance of a Mk3 Dongle Terminator.
	mk3DT := newMk3DT(dongle)
	defer mk3DT.Close()

	// Process CLI Options.
	if raw != "" {
		sendRawCommand(mk3DT, raw)
		return
	}
	if scanCells {
		os.Exit(scanForCells(mk3DT, maxAddr))
	}
	if scanCards {
		os.Exit(scanForCards(mk3DT, maxAddr))
	}
	if newAddr > 0 {
		os.Exit(setAddr(mk3DT, newAddr))
		return
	}
	if realtime {
		os.Exit(listRealtimeValues(mk3DT, maxAddr))
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

// Prints to standard out the result.
func sendRawCommand(mk3DT *Mk3DT, s string) {
	log.Println(string(mk3DT.Raw(s + "\n\r")))
}

// Prints to standard out the result.
func scanForCells(mk3DT *Mk3DT, maxAddr int) int {
	pos := 1
	lastSn := 0
	lastEmpty := false
	for addr := 1; addr <= maxAddr; addr++ {
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

func scanForCards(mk3DT *Mk3DT, maxAddr int) int {
	lastSn := 0
	count := 1
	for addr := 1; addr <= maxAddr; addr++ {
		sn := mk3DT.GetSerialNum(addr)
		if sn != lastSn && sn != 0 {
			cells := mk3DT.GetNumCells(addr)
			fmt.Printf("%02d: S/N %05d with cells %03d-%03d/%d\n\r", count, sn, addr, addr+cells-1, cells)
			lastSn = sn
			count++
		}
	}
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

func listRealtimeValues(mk3DT *Mk3DT, maxAddr int) int {
	for {
		fmt.Printf("\n\r")
		fmt.Printf("|------|-------|------|-------|-------|\n\r")
		fmt.Printf("| CELL | VOLTS | TEMP |  S/N  |  CCC  |\n\r")
		for addr := 1; addr <= maxAddr; addr++ {
			fmt.Printf("|------|-------|------|-------|-------|\n\r")
			if v := mk3DT.GetRealTimeVoltage(addr); v > 0 {
				fmt.Printf("| % 4d ", addr)
				fmt.Printf("| % 5.2f ", v/float32(mk3DT.GetNumCells(addr)))
				fmt.Printf("| % 4d ", mk3DT.GetCellsTemp(addr))
				fmt.Printf("| %05d ", mk3DT.GetSerialNum(addr))
				fmt.Printf("| % 5d ", mk3DT.GetNumCells(addr))
				fmt.Printf("|\n\r")
			} else {
				fmt.Printf("| % 4d |       |      |       |       |\n\r", addr)
			}
		}
		fmt.Printf("|------|-------|------|-------|-------|\n\r")
	}
	return 0
}
