package main

import (
	"bytes"
	"log"
	"strconv"
)

type Mk3DT struct {
	serialPort SerialPort
}

type Mk3DTStatus struct {
	Type         string
	Version      string
	Unit         string
	SerialNumber string
}

type Mk3DTCommands struct {
	Mk3DTStatus
	SetStopChargeTemp         bool
	GetStopTemp               bool
	DisableStopChargeTemp     bool
	ChangeAddr                bool
	DisableShunt              bool
	EnableShunt               bool
	ForceFan                  bool
	GetFirstPosition          bool
	SetFirstPosition          bool
	GetHighVoltage            bool
	ClearMaxVoltageHistory    bool
	ClearMinVoltageHistory    bool
	ClearVoltageHistory       bool
	TriggerLights             bool
	GetMaxVoltage             bool
	GetMinVoltage             bool
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

type Mk3DTLightsStatus struct {
	OverTempRegulator bool
	OverTempCellAddr  int
	MinVoltageSeen    bool
	MaxVoltage        bool
	MinVoltage        bool
	HighVoltage       bool
	ShuntDisabled     bool
}

type Mk3DTResponse struct {
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

func (this *Mk3DT) execCmd(addr int, cmd string, value string) Mk3DTResponse {
	// Clear the Dongle Terminator buffer.
	this.readBytes(0)
	// Send the command.
	this.writeBytes([]byte(strconv.Itoa(addr) + cmd + "." + value + "\n\r"))
	// Read and return a Mk3DTResponse.
	buf := this.readBytes(0)
	cur := 0
	addrNum := []byte{}
	r := Mk3DTResponse{}
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
func (this *Mk3DT) SetStopChargeTemp(addr int, temp int) bool {
	if temp < 32 || temp > 180 {
		return false
	}
	r := this.execCmd(addr, "bt", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return tempToInt(r.Value) == temp
}

func (this *Mk3DT) GetStopTemp(addr int) int {
	r := this.execCmd(addr, "bt", "")
	// Return value as an int.
	return tempToInt(r.Value)
}

func (this *Mk3DT) DisableStopChargeTemp(addr int) bool {
	r := this.execCmd(addr, "btd", "")
	// Check that the returned value equals "DISABLE".
	return r.Value == "DISABLE"
}

// addr 0-255
func (this *Mk3DT) ChangeAddr(addr int, newAddr int) bool {
	r := this.execCmd(addr, "ch", strconv.Itoa(newAddr))
	// Check that the returned value is the same as the sent addr.
	n, _ := strconv.ParseInt(r.Value, 10, 32)
	return newAddr == int(n)
}

func (this *Mk3DT) GetCommands(addr int) Mk3DTCommands {
	this.execCmd(addr, "", "")
	// Return value as Mk3DTCommands.
	return Mk3DTCommands{}
}

func (this *Mk3DT) DisableShunt(addr int) bool {
	r := this.execCmd(addr, "d", "")
	// Check that the returned value equals "Disable".
	return r.Command == "Disable"
}

func (this *Mk3DT) EnableShunt(addr int) bool {
	r := this.execCmd(addr, "e", "")
	// Check that the returned value equals "Enable".
	return r.Command == "Enable"
}

// level 0-8
func (this *Mk3DT) ForceFan(addr int, level int) bool {
	if level < 0 || level > 8 {
		return false
	}
	r := this.execCmd(addr, "f", strconv.Itoa(level))
	// Check that the returned value is the same as the sent level.
	l, _ := strconv.ParseInt(r.Value, 10, 32)
	return level == int(l)
}

func (this *Mk3DT) GetFirstPosition(addr int) bool {
	r := this.execCmd(addr, "fi", "")
	// Return the value as a bool.
	return r.Value == "0"
}

func (this *Mk3DT) SetFirstPosition(addr int, value bool) bool {
	v := "0"
	if value {
		v = "1"
	}
	r := this.execCmd(addr, "fi", v)
	// Check that the returned value is the same as the sent value.
	return r.Value == v
}

func (this *Mk3DT) GetHighVoltage(addr int) float32 {
	r := this.execCmd(addr, "g", "")
	// Return the value as float32.
	return voltToFloat32(r.Value)
}

func (this *Mk3DT) ClearMaxVoltageHistory(addr int) bool {
	this.execCmd(addr, "hma", "")
	// Nothing to return.
	return true
}

func (this *Mk3DT) ClearMinVoltageHistory(addr int) bool {
	this.execCmd(addr, "hmi", "")
	// Nothing to return.
	return true
}

func (this *Mk3DT) ClearVoltageHistory(addr int) bool {
	this.execCmd(addr, "h", "")
	// Nothing to return.
	return true
}

// todo
func (this *Mk3DT) TriggerLights(addr int) Mk3DTLightsStatus {
	this.execCmd(addr, "l", "")
	// Return value as LightsMk3DTStatus.
	return Mk3DTLightsStatus{}
}

func (this *Mk3DT) GetMaxVoltage(addr int) float32 {
	r := this.execCmd(addr, "ma", "")
	// Return the value as float32.
	return voltToFloat32(r.Value)
}

func (this *Mk3DT) GetMinVoltage(addr int) float32 {
	r := this.execCmd(addr, "mi", "")
	// Return the value as float32.
	return voltToFloat32(r.Value)
}

func (this *Mk3DT) GetStopChargeUnderVoltage(addr int) bool {
	r := this.execCmd(addr, "p", "")
	// Return the value as bool.
	return r.Value == "1"
}

func (this *Mk3DT) SetStopChargeUnderVoltage(addr int, stop bool) bool {
	v := "0"
	if stop {
		v = "1"
	}
	r := this.execCmd(addr, "p", v)
	// Return the value as bool.
	return r.Value == "1"
}

func (this *Mk3DT) GetRealTimeVoltage(addr int) float32 {
	r := this.execCmd(addr, "q", "")
	// Return the value as float32.
	return voltToFloat32(r.Value)
}

func (this *Mk3DT) GetLowVoltage(addr int) float32 {
	r := this.execCmd(addr, "r", "")
	// Return the value as float32.
	return voltToFloat32(r.Value)
}

// volts 0.000-9.999
func (this *Mk3DT) SetMaxVoltage(addr int, volts float32) bool {
	if volts < 0.000 || volts > 9.999 {
		return false
	}
	r := this.execCmd(addr, "seth", strconv.FormatFloat(float64(volts), 'f', 3, 32))
	// Check that the returned value is the same as the sent volts.
	return voltToFloat32(r.Value) == volts
}

// volts 0.000-9.999
func (this *Mk3DT) SetMinVoltage(addr int, volts float32) bool {
	if volts < 0.000 || volts > 9.999 {
		return false
	}
	r := this.execCmd(addr, "setl", strconv.FormatFloat(float64(volts), 'f', 3, 32))
	// Check that the returned value is the same as the sent volts.
	return voltToFloat32(r.Value) == volts
}

// volts 0.000-9.999
func (this *Mk3DT) SetOverVoltage(addr int, volts float32) bool {
	if volts < 0.000 || volts > 9.999 {
		return false
	}
	r := this.execCmd(addr, "seto", strconv.FormatFloat(float64(volts), 'f', 3, 32))
	// Check that the returned value is the same as the sent volts.
	return voltToFloat32(r.Value) == volts
}

func (this *Mk3DT) GetStatus(addr int) Mk3DTStatus {
	this.execCmd(addr, "s", "")
	// Return the value as Mk3DTStatus.
	return Mk3DTStatus{}
}

func (this *Mk3DT) GetAddrTemp(addr int) int {
	r := this.execCmd(addr, "t", "")
	// Return the value as int.
	return tempToInt(r.Value)
}

// temp 32-181
func (this *Mk3DT) SetFanMaxTemp(addr int, temp int) bool {
	if temp < 32 || temp > 180 {
		return false
	}
	r := this.execCmd(addr, "temph", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return tempToInt(r.Value) == temp
}

// temp 32-181
func (this *Mk3DT) SetStopDissipatingTemp(addr int, temp int) bool {
	if temp < 32 || temp > 180 {
		return false
	}
	r := this.execCmd(addr, "tempo", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return tempToInt(r.Value) == temp
}

// temp 32-181
func (this *Mk3DT) SetFanLowTemp(addr int, temp int) bool {
	if temp < 32 || temp > 180 {
		return false
	}
	r := this.execCmd(addr, "tempw", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return tempToInt(r.Value) == temp
}

func (this *Mk3DT) GetCellsTemp(addr int) int {
	r := this.execCmd(addr, "x", "")
	// Return the value as int (or string)?
	return tempToInt(r.Value)
}
