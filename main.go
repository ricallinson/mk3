package main

import (
	"flag"
	"github.com/tarm/serial"
	"log"
	"time"
	"bytes"
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
	writeBytes(serialPort, []byte("s.\n\r"))
	log.Println(string(readBytes(serialPort, 0)))
}

func readBytes(serialPort SerialPort, delim byte) []byte {
	limit := 1000
	buff := make([]byte, 1)
	data := make([]byte, 0)
	for limit > 0 {
		i, _ := serialPort.Read(buff)
		// log.Println(i, buff[0], delim)
		if buff[0] == delim || i == 0 {
			break
		}
		data = append(data, buff[0])
		// log.Println(i, string(buff[0]))
		// log.Println(string(data))
		limit--
	}
	data = bytes.TrimSpace(data)
	// log.Println(string(data))
	return data
}

func writeBytes(serialPort SerialPort, b []byte) {
	_, e := serialPort.Write(b)
	if e != nil {
		log.Println(e)
	}
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
