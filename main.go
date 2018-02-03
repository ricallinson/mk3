package main

import (
	"flag"
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
	var dongle string
	flag.StringVar(&dongle, "dongle", "", "Serial port that's connected to the Dongle Terminator.")
	var commands string
	flag.StringVar(&commands, "exe", "", "Path to the YAML configuration file to execute.")
	var addr int
	flag.IntVar(&addr, "addr", -1, "The address to which the commands are to be executed. Defult is all.")
	flag.Parse()
	if commands == "" {
		log.Println("You must provide a path to YAML file with the commands to execute")
		return
	}
	if addr < -1 || addr > 255 {
		log.Println("The address must be between 0 and 255.")
		return
	}
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
	e := NewExecutor(NewMk3DT(serialPort))
	e.Commands = readYamlFileToExecutorCommands(commands)
	// e.ExecuteCommands()
	// e.ExecuteCommandsAtAddr()
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
