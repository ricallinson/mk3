package main

import (
	"flag"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"time"
)

type SerialPort interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Flush() error
	Close() error
}

func main() {
	var raw string
	flag.StringVar(&raw, "raw", "", "Send a raw command to the Dongle Terminator.")
	var scan bool
	flag.BoolVar(&scan, "scan", false, "Sacn the bus for addresses.")
	var dongle string
	flag.StringVar(&dongle, "dongle", "", "Serial port that's connected to the Dongle Terminator.")
	var commands string
	flag.StringVar(&commands, "cmd", "", "Path to the YAML configuration file of commands to execute.")
	var addr int
	flag.IntVar(&addr, "addr", -1, "The address to which the commands are to be executed. Defult is all.")
	var highAddr int
	flag.IntVar(&highAddr, "highaddr", 255, "The highest address to which the commands are to be executed. Defult is 255.")
	flag.Parse()
	var serialPort SerialPort
	var serialError error
	if dongle == "" {
		serialPort = NewMockPort()
	} else {
		serialError, serialPort = connectToDongle(dongle)
	}
	if serialError != nil {
		log.Print("Error opening port: ")
		log.Println(serialError)
		return
	}
	mk3DT := NewMk3DT(serialPort)
	if raw != "" {
		log.Println(string(mk3DT.Raw(raw + "\n\r")))
		return
	}
	if scan {
		for addr := 0; addr <= 255; addr++ {
			if mk3DT.GetStopChargeTemp(addr) > 0 {
				fmt.Println(addr)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("Done")
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
	e := NewExecutor(mk3DT)
	e.Commands = readYamlFileToExecutorCommands(commands)
	var yaml []byte
	if addr >= 0 && addr <= 255 {
		yaml = interfaceToYaml(e.ExecuteCommandsAtAddr(addr))
	} else {
		yaml = interfaceToYaml(e.ExecuteCommands(highAddr))
	}
	log.Println(string(yaml))
	e.Close()
}

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
