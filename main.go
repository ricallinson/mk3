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
	flag.StringVar(&dongle, "dongle", "", "Serial port that's connected to the Dongle Terminator")
	flag.Parse()
	var serialPort SerialPort
	var serialError error
	if dongle == "" {
		log.Println("You must provide the serial port the Dongle Terminator is connected to.")
		return
	} else {
		serialError, serialPort = connectToDongle(dongle)
	}
	if serialError != nil {
		log.Println(serialError)
		return
	}
	mk3DT := NewMk3DT(serialPort)
	mk3DT.GetMaxVolage(1)
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
