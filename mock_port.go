package main

import (
	"bytes"
	// "log"
	"strconv"
	"strings"
)

type MockPort struct {
	buffer []byte
}

type mk3DTRequest struct {
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
	r := &mk3DTRequest{}
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

// Checks the mk3DTRequest and updates the buffer if needed.
func (this *MockPort) processRequest(r *mk3DTRequest) {
	// log.Print("Mock Request Recived:")
	// log.Println(r)
	switch {
	case strings.HasPrefix(r.Command, "btd"):
		this.DisableStopChargeTemp(r)
	case strings.HasPrefix(r.Command, "bt"):
		this.GetSetStopChargeTemp(r)
	case strings.HasPrefix(r.Command, "ch"):
		this.ChangeAddr(r)
	case r.Command == "":
		this.GetCommands(r)
	case strings.HasPrefix(r.Command, "d"):
		this.DisableShunt(r)
	case strings.HasPrefix(r.Command, "e"):
		this.EnableShunt(r)
	case strings.HasPrefix(r.Command, "fi"):
		this.GetSetFirstPosition(r)
	case strings.HasPrefix(r.Command, "f"):
		this.ForceFan(r)
	case strings.HasPrefix(r.Command, "g"):
		this.GetHighVoltage(r)
	case strings.HasPrefix(r.Command, "hma"):
		this.ClearMaxVoltageHistory(r)
	case strings.HasPrefix(r.Command, "hmi"):
		this.ClearMinVoltageHistory(r)
	case strings.HasPrefix(r.Command, "h"):
		this.ClearVoltageHistory(r)
	case strings.HasPrefix(r.Command, "l"):
		this.TriggerLights(r)
	case strings.HasPrefix(r.Command, "ma"):
		this.GetMaxVoltage(r)
	case strings.HasPrefix(r.Command, "mi"):
		this.GetMinVoltage(r)
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
	case strings.HasPrefix(r.Command, "temph"):
		this.SetFanMaxTemp(r)
	case strings.HasPrefix(r.Command, "tempo"):
		this.SetStopDissipatingTemp(r)
	case strings.HasPrefix(r.Command, "tempw"):
		this.SetFanLowTemp(r)
	case strings.HasPrefix(r.Command, "t"):
		this.GetAddrTemp(r)
	case strings.HasPrefix(r.Command, "x"):
		this.GetCellsTemp(r)
	}
}

func (this *MockPort) bufferResponse(addr int, value string) {
	this.buffer = append(this.buffer, []byte(padInt(addr)+value)...)
	// log.Println(string(this.buffer))
}

// temp 32-180 F
func (this *MockPort) GetSetStopChargeTemp(r *mk3DTRequest) {
	if r.Value == "" {
		this.bufferResponse(r.Addr, "BT 180F")
	} else {
		// Check that the returned value is the same as the sent temp.
		this.bufferResponse(r.Addr, "BT "+r.Value+"F")
	}
}

func (this *MockPort) DisableStopChargeTemp(r *mk3DTRequest) {
	// Check that the returned value equals "DISABLE".
	this.bufferResponse(r.Addr, "BT DISABLE")
}

// addr 0-255
func (this *MockPort) ChangeAddr(r *mk3DTRequest) {
	// Check that the returned value is the same as the sent addr.
	this.bufferResponse(r.Addr, "00 Now:"+r.Value)
}

func (this *MockPort) GetCommands(r *mk3DTRequest) {
	this.bufferResponse(r.Addr, `--RUDMAN MK3X8 REGULATOR
--V1.28 UNIT:09-12 S/N: 00495
COMMANDL VOLTAGE  CHANGEAD DISABLE  
ENABLE   FAN      GETHIGHV HSTCLEAR 
HMACLEAR HMICLEAR LIGHTS   MINVOLTS 
MAXVOLTS PHEV     QUERYTOT READLOWV 
STATUS   SETHIGH  SETLOW   SETOVER  
TEMPERAT TEMPWARM TEMPHOT  TEMPOFF  
XTRNTEMP BTEMP    BTDISABL FIRSTPOS`)
}

func (this *MockPort) DisableShunt(r *mk3DTRequest) {
	// Check that the returned value equals "Disable".
	this.bufferResponse(r.Addr, "Disable")
}

func (this *MockPort) EnableShunt(r *mk3DTRequest) {
	// Check that the returned value equals "Enable".
	this.bufferResponse(r.Addr, "Enable")
}

// level 0-8
func (this *MockPort) ForceFan(r *mk3DTRequest) {
	// Check that the returned value is the same as the sent level.
	this.bufferResponse(r.Addr, "F "+r.Value)
}

func (this *MockPort) GetSetFirstPosition(r *mk3DTRequest) {
	if r.Value == "" {
		// Return the value as a bool.
		this.bufferResponse(r.Addr, "FP 0")
		return
	}
	v := "0"
	if r.Value == "1" {
		v = "1"
	}
	// Check that the returned value is the same as the sent value.
	this.bufferResponse(r.Addr, "FP "+v)
}

func (this *MockPort) GetHighVoltage(r *mk3DTRequest) {
	// Return the value as float32.
	this.bufferResponse(r.Addr, "G 3.9V")
}

func (this *MockPort) ClearMaxVoltageHistory(r *mk3DTRequest) {
	// Nothing to return.
}

func (this *MockPort) ClearMinVoltageHistory(r *mk3DTRequest) {
	// Nothing to return.
}

func (this *MockPort) ClearVoltageHistory(r *mk3DTRequest) {
	// Nothing to return.
}

func (this *MockPort) TriggerLights(r *mk3DTRequest) {
	// Return value as LightsStatus.
}

func (this *MockPort) GetMaxVoltage(r *mk3DTRequest) {
	// Return the value as float32.
	this.bufferResponse(r.Addr, "MA 3.971V")
}

func (this *MockPort) GetMinVoltage(r *mk3DTRequest) {
	// Return the value as float32.
	this.bufferResponse(r.Addr, "MA 2.432V")
}

func (this *MockPort) GetSetStopChargeUnderVoltage(r *mk3DTRequest) {
	if r.Value == "" {
		// Return the value as a bool.
		this.bufferResponse(r.Addr, "P 0")
		return
	}
	v := "0"
	if r.Value == "1" {
		v = "1"
	}
	// Return the value as bool.
	this.bufferResponse(r.Addr, "P "+v)
}

func (this *MockPort) GetRealTimeVoltage(r *mk3DTRequest) {
	// Return the value as float32.
	this.bufferResponse(r.Addr, "Q 3.4V")
}

func (this *MockPort) GetLowVoltage(r *mk3DTRequest) {
	// Return the value as float32.
	this.bufferResponse(r.Addr, "R 2.432V")
}

// volts 0.000-9.999
func (this *MockPort) SetMaxVoltage(r *mk3DTRequest) {
	// Check that the returned value is the same as the sent volts.
	this.bufferResponse(r.Addr, "H "+r.Value+"V")
}

// volts 0.000-9.999
func (this *MockPort) SetMinVoltage(r *mk3DTRequest) {
	// Check that the returned value is the same as the sent volts.
	this.bufferResponse(r.Addr, "H "+r.Value+"V")
}

// volts 0.000-9.999
func (this *MockPort) SetOverVoltage(r *mk3DTRequest) {
	// Check that the returned value is the same as the sent volts.
	this.bufferResponse(r.Addr, "H "+r.Value+"V")
}

func (this *MockPort) GetStatus(r *mk3DTRequest) {
	// Return the value as Status.
}

func (this *MockPort) GetAddrTemp(r *mk3DTRequest) {
	// Return the value as int.
	this.bufferResponse(r.Addr, "T 120F")
}

// temp 32-181
func (this *MockPort) SetFanMaxTemp(r *mk3DTRequest) {
	// Check that the returned value is the same as the sent temp.
	this.bufferResponse(r.Addr, "TH "+r.Value+"F")
}

// temp 32-181
func (this *MockPort) SetStopDissipatingTemp(r *mk3DTRequest) {
	// Check that the returned value is the same as the sent temp.
	this.bufferResponse(r.Addr, "TO "+r.Value+"F")
}

// temp 32-181
func (this *MockPort) SetFanLowTemp(r *mk3DTRequest) {
	// Check that the returned value is the same as the sent temp.
	this.bufferResponse(r.Addr, "W "+r.Value+"F")
}

func (this *MockPort) GetCellsTemp(r *mk3DTRequest) {
	// Return the value as int (or string)?
	this.bufferResponse(r.Addr, "X Cold")
}
