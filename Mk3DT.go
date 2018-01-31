package main

import (
	"bytes"
	"log"
	"strconv"
	"strings"
)

type Mk3DT struct {
	serialPort SerialPort
}

type Status struct {
	Type         string
	Version      string
	Unit         string
	SerialNumber string
}

type Commands struct {
	Status
	SetStopTemp               bool
	GetStopTemp               bool
	DisableStopTemp           bool
	ChangeAddr                bool
	DisableShunt              bool
	EnableShunt               bool
	ForceFan                  bool
	GetFirstPosition          bool
	SetFirstPosition          bool
	GetHighVoltage            bool
	ClearMaxVolageHistory     bool
	ClearMinVolageHistory     bool
	ClearVolageHistory        bool
	TriggerLights             bool
	GetMaxVolage              bool
	GetMinVolage              bool
	SetStopChargeUnderVoltage bool
	GetRealTimeVoltage        bool
	GetLowVoltage             bool
	SetMaxVoltage             bool
	SetMinVoltage             bool
	SetOverVoltage            bool
	GetStatus                 bool
	GetAddrTemp               bool
	SetFanMaxTemp             bool
	SetStopDissipatingTemp    bool
	SetFanLowTemp             bool
	GetCellsTemp              bool
}

type LightsStatus struct {
	OverTempRegulator bool
	OverTempCellAddr  int
	MinVoltageSeen    bool
	MaxVoltage        bool
	MinVoltage        bool
	HighVoltage       bool
	ShuntDisabled     bool
}

type Response struct {
	Addr    int
	Command string
	Value   string
}

func NewMk3DT(p SerialPort) *Mk3DT {
	this := &Mk3DT{
		serialPort: p,
	}
	return this
}

func (this *Mk3DT) readBytes(delim byte) []byte {
	limit := 100
	buff := make([]byte, 1)
	data := make([]byte, 0)
	for limit > 0 {
		i, _ := this.serialPort.Read(buff)
		if i == 0 {
			break
		}
		data = append(data, buff[0])
		limit--
	}
	data = bytes.TrimSpace(data)
	return data
}

func (this *Mk3DT) writeBytes(b []byte) {
	_, e := this.serialPort.Write(b)
	if e != nil {
		log.Println(e)
	}
}

func (this *Mk3DT) execCmd(addr int, cmd string, value string) Response {
	// Clear the Dongle Terminator buffer.
	this.readBytes(0)
	// Send the command.
	this.writeBytes([]byte(strconv.Itoa(addr) + cmd + "." + value + "\n\r"))
	// Read and return a Response.
	buf := this.readBytes(0)
	cur := 0
	addrNum := []byte{}
	r := Response{}
	for cur < len(buf) {
		if buf[cur] >= 48 && buf[cur] <= 57 {
			// If 0-9, append to Addr.
			addrNum = append(addrNum, buf[cur])
		} else if buf[cur] >= 65 && buf[cur] <= 122 {
			// Id a-zA-Z, append to Command.
			r.Command += string(buf[cur])
		} else if buf[cur] == 32 || buf[cur] == 45 {
			// After a space or - append to Value.
			r.Value = string(bytes.TrimSpace(buf[cur+1:]))
			break
		}
		cur++
	}
	r.Addr, _ = strconv.Atoi(string(addrNum))
	return r
}

// temp 32-180 F
func (this *Mk3DT) SetStopTemp(addr int, temp int) bool {
	r := this.execCmd(addr, "bt", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	t, _ := strconv.ParseInt(strings.TrimSuffix(r.Value, "F"), 10, 32)
	// log.Println(temp, r.Value, t)
	return int(t) == temp
}

func (this *Mk3DT) GetStopTemp(addr int) int {
	r := this.execCmd(addr, "bt", "")
	// Return value as an int.
	t, _ := strconv.ParseInt(strings.TrimSuffix(r.Value, "F"), 10, 32)
	return int(t)
}

func (this *Mk3DT) DisableStopTemp(addr int) bool {
	r := this.execCmd(addr, "btd", "")
	// Check that the returned value equals "DISABLE".
	// log.Print(r)
	return r.Value == "DISABLE"
}

// addr 0-255
func (this *Mk3DT) ChangeAddr(addr int, newAddr int) bool {
	r := this.execCmd(addr, "ch", strconv.Itoa(newAddr))
	// Check that the returned value is the same as the sent addr.
	n, _ := strconv.ParseInt(r.Value, 10, 32)
	return newAddr == int(n)
}

func (this *Mk3DT) GetCommands(addr int) Commands {
	this.execCmd(addr, "", "")
	// Return value as Commands.
	return Commands{}
}

func (this *Mk3DT) DisableShunt(addr int) bool {
	this.execCmd(addr, "d", "")
	// Check that the returned value equals "Disable".
	return false
}

func (this *Mk3DT) EnableShunt(addr int) bool {
	this.execCmd(addr, "e", "")
	// Check that the returned value equals "Enable".
	return false
}

// level 0-8
func (this *Mk3DT) ForceFan(addr int, level int) bool {
	this.execCmd(addr, "f", strconv.Itoa(level))
	// Check that the returned value is the same as the sent level.
	return false
}

func (this *Mk3DT) GetFirstPosition(addr int) bool {
	this.execCmd(addr, "fi", "")
	// Return the value as a bool.
	return false
}

func (this *Mk3DT) SetFirstPosition(addr int, value bool) bool {
	v := "0"
	if value {
		v = "1"
	}
	this.execCmd(addr, "fi", v)
	// Check that the returned value is the same as the sent value.
	return false
}

func (this *Mk3DT) GetHighVoltage(addr int) float32 {
	this.execCmd(addr, "g", "")
	// Return the value as float32.
	return 0.0
}

func (this *Mk3DT) ClearMaxVolageHistory(addr int) {
	this.execCmd(addr, "hma", "")
	// Nothing to return.
}

func (this *Mk3DT) ClearMinVolageHistory(addr int) {
	this.execCmd(addr, "hmi", "")
	// Nothing to return.
}

func (this *Mk3DT) ClearVolageHistory(addr int) {
	this.execCmd(addr, "h", "")
	// Nothing to return.
}

func (this *Mk3DT) TriggerLights(addr int) LightsStatus {
	this.execCmd(addr, "l", "")
	// Return value as LightsStatus.
	return LightsStatus{}
}

func (this *Mk3DT) GetMaxVolage(addr int) float32 {
	r := this.execCmd(addr, "ma", "")
	// Return the value as float32.
	f, _ := strconv.ParseFloat(strings.TrimSuffix(r.Value, "V"), 32)
	return float32(f)
}

func (this *Mk3DT) GetMinVolage(addr int) float32 {
	this.execCmd(addr, "mi", "")
	// Return the value as float32.
	return 0.0
}

func (this *Mk3DT) GetStopChargeUnderVoltage(addr int) float32 {
	this.execCmd(addr, "p", "")
	// Return the value as float32.
	return 0.0
}

func (this *Mk3DT) SetStopChargeUnderVoltage(addr int, stop bool) float32 {
	this.execCmd(addr, "p", strconv.FormatBool(stop))
	// Return the value as float32.
	return 0.0
}

func (this *Mk3DT) GetRealTimeVoltage(addr int) float32 {
	this.execCmd(addr, "q", "")
	// Return the value as float32.
	return 0.0
}

func (this *Mk3DT) GetLowVoltage(addr int) float32 {
	this.execCmd(addr, "r", "")
	// Return the value as float32.
	return 0.0
}

// volts 0.000-9.999
func (this *Mk3DT) SetMaxVoltage(addr int, volts int) bool {
	this.execCmd(addr, "seth", strconv.Itoa(volts))
	// Check that the returned value is the same as the sent volts.
	return false
}

// volts 0.000-9.999
func (this *Mk3DT) SetMinVoltage(addr int, volts int) bool {
	this.execCmd(addr, "setl", strconv.Itoa(volts))
	// Check that the returned value is the same as the sent volts.
	return false
}

// volts 0.000-9.999
func (this *Mk3DT) SetOverVoltage(addr int, volts int) bool {
	this.execCmd(addr, "seto", strconv.Itoa(volts))
	// Check that the returned value is the same as the sent volts.
	return false
}

func (this *Mk3DT) GetStatus(addr int) Status {
	this.execCmd(addr, "s", "")
	// Return the value as Status.
	return Status{}
}

func (this *Mk3DT) GetAddrTemp(addr int) int {
	this.execCmd(addr, "t", "")
	// Return the value as int.
	return 0
}

// temp 32-181
func (this *Mk3DT) SetFanMaxTemp(addr int, temp int) bool {
	this.execCmd(addr, "temph", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return false
}

// temp 32-181
func (this *Mk3DT) SetStopDissipatingTemp(addr int, temp int) bool {
	this.execCmd(addr, "tempo", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return false
}

// temp 32-181
func (this *Mk3DT) SetFanLowTemp(addr int, temp int) bool {
	this.execCmd(addr, "tempw", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return false
}

func (this *Mk3DT) GetCellsTemp(addr int) int {
	this.execCmd(addr, "x", "")
	// Return the value as int (or string)?
	return 0
}
