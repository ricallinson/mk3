package main

import (
	"bytes"
	"log"
	"strconv"
)

type Mk3 struct {
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

func NewMk3(p SerialPort) *Mk3 {
	this := &Mk3{
		serialPort: p,
	}
	return this
}

func (this *Mk3) readBytes(delim byte) []byte {
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

func (this *Mk3) writeBytes(b []byte) {
	_, e := this.serialPort.Write(b)
	if e != nil {
		log.Println(e)
	}
}

func (this *Mk3) execCmd(addr int, cmd string, value string) string {
	// Clear the Dongle Terminator buffer.
	this.readBytes(0)
	// Send the command.
	this.writeBytes([]byte(strconv.Itoa(addr) + cmd + "." + value + "\n\r"))
	// Read and return the response.
	return string(this.readBytes(0))
}

// temp 32-180 F
func (this *Mk3) SetStopTemp(addr int, temp int) bool {
	this.execCmd(addr, "bt", strconv.Itoa(temp))
	// Check that the returned temp is the same as the sent temp.
	return false
}

func (this *Mk3) GetStopTemp(addr int) int {
	this.execCmd(addr, "bt", "")
	// Returned the given temp as an int.
	return 0
}

func (this *Mk3) DisableStopTemp(addr int) bool {
	this.execCmd(addr, "btd", "")
	// Check that the returned value equals "DISABLE".
	return false
}

// addr 0-255
func (this *Mk3) ChangeAddr(addr int, newAddr int) bool {
	this.execCmd(addr, "ch", strconv.Itoa(newAddr))
	// Check that the returned addr is the same as the sent addr.
	return false
}

func (this *Mk3) GetCommands(addr int) Commands {
	this.execCmd(addr, "", "")
	// Retrun commands listed as a Struct.
	return Commands{}
}

func (this *Mk3) DisableShunt(addr int) bool {
	this.execCmd(addr, "d", "")
	// Check that the returned value equals "Disable".
	return false
}

func (this *Mk3) EnableShunt(addr int) bool {
	this.execCmd(addr, "e", "")
	// Check that the returned value equals "Enable".
	return false
}

// level 0-8
func (this *Mk3) ForceFan(addr int, level int) bool {
	this.execCmd(addr, "f", strconv.Itoa(level))
	// Check that the returned level is the same as the sent level.
	return false
}

func (this *Mk3) GetFirstPosition(addr int) bool {
	this.execCmd(addr, "fi", "")
	// Return the value as a bool.
	return false
}

func (this *Mk3) SetFirstPosition(addr int, value bool) bool {
	v := "0"
	if value {
		v = "1"
	}
	this.execCmd(addr, "fi", v)
	// Check that the returned value is the same as the sent value.
	return false
}

func (this *Mk3) GetHighVoltage(addr int) float32 {
	this.execCmd(addr, "g", "")
	// Return the voltage as float.
	return 0.0
}

func (this *Mk3) ClearMaxVolageHistory(addr int) {
	this.execCmd(addr, "hma", "")
}

func (this *Mk3) ClearMinVolageHistory(addr int) {
	this.execCmd(addr, "hmi", "")
}

func (this *Mk3) ClearVolageHistory(addr int) {
	this.execCmd(addr, "h", "")
}

func (this *Mk3) TriggerLights(addr int) LightsStatus {
	this.execCmd(addr, "l", "")
	// Return data as a Struct.
	return LightsStatus{}
}

func (this *Mk3) GetMaxVolage(addr int) float32 {
	this.execCmd(addr, "ma", "")
	// Return the voltage as float32.
	return 0.0
}

func (this *Mk3) GetMinVolage(addr int) float32 {
	this.execCmd(addr, "mi", "")
	// Return the voltage as float32.
	return 0.0
}

func (this *Mk3) SetStopChargeUnderVoltage(addr int, stop bool) float32 {
	this.execCmd(addr, "p", strconv.FormatBool(stop))
	// Return the voltage as float32.
	return 0.0
}

func (this *Mk3) GetRealTimeVoltage(addr int) float32 {
	this.execCmd(addr, "q", "")
	// Return the voltage as float32.
	return 0.0
}

func (this *Mk3) GetLowVoltage(addr int) float32 {
	this.execCmd(addr, "r", "")
	// Return the voltage as float32.
	return 0.0
}

// volts 0.000-9.999
func (this *Mk3) SetMaxVoltage(addr int, volts int) bool {
	this.execCmd(addr, "seth", strconv.Itoa(volts))
	// Check that the returned volts is the same as the sent volts.
	return false
}

// volts 0.000-9.999
func (this *Mk3) SetMinVoltage(addr int, volts int) bool {
	this.execCmd(addr, "setl", strconv.Itoa(volts))
	// Check that the returned volts is the same as the sent volts.
	return false
}

// volts 0.000-9.999
func (this *Mk3) SetOverVoltage(addr int, volts int) bool {
	this.execCmd(addr, "seto", strconv.Itoa(volts))
	// Check that the returned volts is the same as the sent volts.
	return false
}

func (this *Mk3) GetStatus(addr int) {
	this.execCmd(addr, "s", "")
	// Return data as a Struct.
}

func (this *Mk3) GetAddrTemp(addr int) int {
	this.execCmd(addr, "t", "")
	// Return the temp as int.
	return 0
}

// temp 32-181
func (this *Mk3) SetFanMaxTemp(addr int, temp int) bool {
	this.execCmd(addr, "temph", strconv.Itoa(temp))
	// Check that the returned temp is the same as the sent temp.
	return false
}

// temp 32-181
func (this *Mk3) SetStopDissipatingTemp(addr int, temp int) bool {
	this.execCmd(addr, "tempo", strconv.Itoa(temp))
	// Check that the returned temp is the same as the sent temp.
	return false
}

// temp 32-181
func (this *Mk3) SetFanLowTemp(addr int, temp int) bool {
	this.execCmd(addr, "tempw", strconv.Itoa(temp))
	// Check that the returned temp is the same as the sent temp.
	return false
}

func (this *Mk3) GetCellsTemp(addr int) int {
	this.execCmd(addr, "x", "")
	// Return the temp as int or string?
	return 0
}
