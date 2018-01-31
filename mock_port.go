package main

import (
	"bytes"
	"log"
	"strconv"
	"strings"
)

type MockPort struct {
	buffer []byte
}

type request struct {
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
	r := &request{}
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
	this.processRequest(r)
	return len(b), nil
}

func (this *MockPort) Flush() error {
	return nil
}

func (this *MockPort) Close() error {
	return nil
}

// Checks the request and updates the buffer if needed.
func (this *MockPort) processRequest(r *request) {
	log.Print("Mock Request: ")
	log.Println(r)
	switch {
	case strings.HasPrefix(r.Command, "bt"):
		this.GetSetStopTemp(r)
	case strings.HasPrefix(r.Command, "btd"):
		this.DisableStopTemp(r)
	case strings.HasPrefix(r.Command, "ch"):
		this.ChangeAddr(r)
	case r.Command == "":
		this.GetCommands(r)
	case strings.HasPrefix(r.Command, "d"):
		this.DisableShunt(r)
	case strings.HasPrefix(r.Command, "e"):
		this.EnableShunt(r)
	case strings.HasPrefix(r.Command, "f"):
		this.ForceFan(r)
	case strings.HasPrefix(r.Command, "fi"):
		this.GetFirstPosition(r)
	case strings.HasPrefix(r.Command, "g"):
		this.GetHighVoltage(r)
	case strings.HasPrefix(r.Command, "hma"):
		this.ClearMaxVolageHistory(r)
	case strings.HasPrefix(r.Command, "hmi"):
		this.ClearMinVolageHistory(r)
	case strings.HasPrefix(r.Command, "h"):
		this.ClearVolageHistory(r)
	case strings.HasPrefix(r.Command, "l"):
		this.TriggerLights(r)
	case strings.HasPrefix(r.Command, "ma"):
		this.GetMaxVolage(r)
	case strings.HasPrefix(r.Command, "mi"):
		this.GetMinVolage(r)
	case strings.HasPrefix(r.Command, "p"):
		this.GetSetStopChargeUnderVoltage(r)
	case strings.HasPrefix(r.Command, "q"):
		this.GetRealTimeVoltage(r)
	case strings.HasPrefix(r.Command, "r"):
		this.GetLowVoltage(r)
	case strings.HasPrefix(r.Command, "seth"):
		this.SetMaxVoltage(r)
	case strings.HasPrefix(r.Command, "setl"):
		this.SetMinVoltage(r)
	case strings.HasPrefix(r.Command, "seto"):
		this.SetOverVoltage(r)
	case strings.HasPrefix(r.Command, "s"):
		this.GetStatus(r)
	case strings.HasPrefix(r.Command, "t"):
		this.GetAddrTemp(r)
	case strings.HasPrefix(r.Command, "temph"):
		this.SetFanMaxTemp(r)
	case strings.HasPrefix(r.Command, "tempo"):
		this.SetStopDissipatingTemp(r)
	case strings.HasPrefix(r.Command, "tempw"):
		this.SetFanLowTemp(r)
	case strings.HasPrefix(r.Command, "x"):
		this.GetCellsTemp(r)
	}
}

// temp 32-180 F
func (this *MockPort) GetSetStopTemp(r *request) {
	// Check that the returned value is the same as the sent temp.
}

func (this *MockPort) DisableStopTemp(r *request) {
	// Check that the returned value equals "DISABLE".
}

// addr 0-255
func (this *MockPort) ChangeAddr(r *request) {
	// Check that the returned value is the same as the sent addr.
}

func (this *MockPort) GetCommands(r *request) {
	// Return value as Commands.
}

func (this *MockPort) DisableShunt(r *request) {
	// Check that the returned value equals "Disable".
}

func (this *MockPort) EnableShunt(r *request) {
	// Check that the returned value equals "Enable".
}

// level 0-8
func (this *MockPort) ForceFan(r *request) {
	// Check that the returned value is the same as the sent level.
}

func (this *MockPort) GetFirstPosition(r *request) {
	// Return the value as a bool.
}

func (this *MockPort) SetFirstPosition(r *request) {
	// Check that the returned value is the same as the sent value.
}

func (this *MockPort) GetHighVoltage(r *request) {
	// Return the value as float32.
}

func (this *MockPort) ClearMaxVolageHistory(r *request) {
	// Nothing to return.
}

func (this *MockPort) ClearMinVolageHistory(r *request) {
	// Nothing to return.
}

func (this *MockPort) ClearVolageHistory(r *request) {
	// Nothing to return.
}

func (this *MockPort) TriggerLights(r *request) {
	// Return value as LightsStatus.
}

func (this *MockPort) GetMaxVolage(r *request) {
	// Return the value as float32.
	this.buffer = append(this.buffer, []byte(padInt(r.Addr)+"MA 3.971V")...)
}

func (this *MockPort) GetMinVolage(r *request) {
	// Return the value as float32.
}

func (this *MockPort) GetSetStopChargeUnderVoltage(r *request) {
	// Return the value as float32.
}

func (this *MockPort) GetRealTimeVoltage(r *request) {
	// Return the value as float32.
}

func (this *MockPort) GetLowVoltage(r *request) {
	// Return the value as float32.
}

// volts 0.000-9.999
func (this *MockPort) SetMaxVoltage(r *request) {
	// Check that the returned value is the same as the sent volts.
}

// volts 0.000-9.999
func (this *MockPort) SetMinVoltage(r *request) {
	// Check that the returned value is the same as the sent volts.
}

// volts 0.000-9.999
func (this *MockPort) SetOverVoltage(r *request) {
	// Check that the returned value is the same as the sent volts.
}

func (this *MockPort) GetStatus(r *request) {
	// Return the value as Status.
}

func (this *MockPort) GetAddrTemp(r *request) {
	// Return the value as int.
}

// temp 32-181
func (this *MockPort) SetFanMaxTemp(r *request) {
	// Check that the returned value is the same as the sent temp.
}

// temp 32-181
func (this *MockPort) SetStopDissipatingTemp(r *request) {
	// Check that the returned value is the same as the sent temp.
}

// temp 32-181
func (this *MockPort) SetFanLowTemp(r *request) {
	// Check that the returned value is the same as the sent temp.
}

func (this *MockPort) GetCellsTemp(r *request) {
	// Return the value as int (or string)?
}
