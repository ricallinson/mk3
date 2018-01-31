package main

import (
	"bytes"
	"log"
	"strconv"
)

type MockPort struct {
	buffer []byte
}

type Request struct {
	Addr    int
	Command string
	Value   string
}

func NewMockPort() *MockPort {
	this := &MockPort{}
	return this
}

func (this *MockPort) Read(b []byte) (int, error) {
	i := copyIntoArray(this.buffer, b)
	this.buffer = this.buffer[i:]
	return i, nil
}

func (this *MockPort) Write(b []byte) (int, error) {
	r := Request{}
	cur := 0
	addr := []byte{}
	for cur < len(b) {
		if b[cur] >= 48 && b[cur] <= 57 {
			// If 0-9, append to Addr.
			addr = append(addr, b[cur])
		} else if b[cur] >= 65 && b[cur] <= 122 {
			// Id A-Z, append to Command.
			r.Command += string(b[cur])
		} else if b[cur] == 46 {
			// After a . append to Value.
			r.Value = string(bytes.TrimSpace(b[cur+1:]))
			break
		}
		cur++
	}
	r.Addr, _ = strconv.Atoi(string(addr))
	log.Println(r)
	return len(b), nil
}

func (this *MockPort) Flush() error {
	return nil
}

func (this *MockPort) Close() error {
	return nil
}
