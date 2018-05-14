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

func (this *Mk3DT) Close() {
	this.serialPort.Close()
}

func (this *Mk3DT) readBytes(delim byte) []byte {
	limit := 1000
	buff := make([]byte, 1)
	data := make([]byte, 0)
	for limit > 0 {
		i, _ := this.serialPort.Read(buff)
		if buff[0] == delim || i == 0 {
			break
		}
		data = append(data, buff[0])
		limit--
	}
	data = bytes.TrimSpace(data)
	return data
}

func (this *Mk3DT) writeBytes(b []byte) int {
	l, e := this.serialPort.Write(b)
	if e != nil {
		log.Println(e)
	}
	// Do not flush after write. // this.serialPort.Flush()
	return l
}

func (this *Mk3DT) execCmd(addr int, cmd string, value string) Mk3DTResponse {
	// Send the command.
	this.writeBytes([]byte(strconv.Itoa(addr) + cmd + "." + value + "\n\r"))
	// time.Sleep(100 * time.Millisecond)
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

func (this *Mk3DT) Raw(c string) []byte {
	this.writeBytes([]byte(c))
	return this.readBytes(0)
}

// temp 32-180 F
func (this *Mk3DT) SetStopChargeTemp(addr int, temp int) bool {
	if temp < 32 || temp > 180 {
		return false
	}
	r := this.execCmd(addr, "btemp", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return tempToInt(r.Value) == temp
}

func (this *Mk3DT) GetStopChargeTemp(addr int) int {
	r := this.execCmd(addr, "btemp", "")
	// Return value as an int.
	return tempToInt(r.Value)
}

func (this *Mk3DT) DisableStopChargeTemp(addr int) bool {
	r := this.execCmd(addr, "btdisable", "")
	// Check that the returned value equals "DISABLE".
	return r.Value == "DISABLE"
}

// addr 0-255
func (this *Mk3DT) ChangeAddr(addr int, newAddr int) bool {
	r := this.execCmd(addr, "changead", strconv.Itoa(newAddr))
	// Check that the returned value is the same as the sent addr.
	if len(r.Value) < 5 {
		return false
	}
	n, _ := strconv.ParseInt(r.Value[4:], 10, 32)
	return newAddr == int(n)
}

func (this *Mk3DT) GetSerialNum(addr int) int {
	r := this.execCmd(addr, "", "")
	// Return the units start and end addresses.
	start := strings.Index(r.Value, " S/N:") + 6
	if len(r.Value) < start+5 {
		return 0
	}
	sn, _ := strconv.Atoi(r.Value[start : start+5])
	return sn
}

func (this *Mk3DT) GetNumCells(addr int) int {
	start, end := this.GetCellRange(addr)
	// Return the units number of cells.
	if start < 1 && end < 1 {
		return 0
	}
	return end - start + 1
}

func (this *Mk3DT) GetCellRange(addr int) (int, int) {
	r := this.execCmd(addr, "", "")
	// Return the units start and end addresses.
	start := strings.Index(r.Value, "UNIT:") + 5
	end := strings.Index(r.Value, " S/N:")
	if start < 0 || end < 0 {
		return 0, 0
	}
	fromTo := strings.Split(strings.TrimSpace(r.Value[start:end]), "-")
	if len(fromTo) != 2 {
		return 0, 0
	}
	from, _ := strconv.Atoi(fromTo[0])
	to, _ := strconv.Atoi(fromTo[1])
	return from, to
}

func (this *Mk3DT) DisableShunt(addr int) bool {
	r := this.execCmd(addr, "disable", "")
	// Check that the returned value equals "Disable".
	return r.Command == "Disable"
}

func (this *Mk3DT) EnableShunt(addr int) bool {
	r := this.execCmd(addr, "enable", "")
	// Check that the returned value equals "Enable".
	return r.Command == "Enable"
}

// level 0-8
func (this *Mk3DT) ForceFan(addr int, level int) bool {
	if level < 0 || level > 8 {
		return false
	}
	r := this.execCmd(addr, "fan", strconv.Itoa(level))
	// Check that the returned value is the same as the sent level.
	l, _ := strconv.ParseInt(r.Value, 10, 32)
	return level == int(l)
}

func (this *Mk3DT) GetFirstPosition(addr int) bool {
	r := this.execCmd(addr, "firstpos", "")
	// Return the value as a bool.
	return r.Value == "1"
}

func (this *Mk3DT) SetFirstPosition(addr int, value int) bool {
	if value < 0 || value > 1 {
		return false
	}
	r := this.execCmd(addr, "firstpos", strconv.Itoa(value))
	// Check that the returned value is the same as the sent value.
	return r.Value == strconv.Itoa(value)
}

func (this *Mk3DT) GetHighVoltage(addr int) float32 {
	r := this.execCmd(addr, "gethighv", "")
	// Return the value as float32.
	return voltToFloat32(r.Value)
}

func (this *Mk3DT) ClearMaxVoltageHistory(addr int) bool {
	this.execCmd(addr, "hmaclear", "")
	// Nothing to return.
	return true
}

func (this *Mk3DT) ClearMinVoltageHistory(addr int) bool {
	this.execCmd(addr, "hmiclear", "")
	// Nothing to return.
	return true
}

func (this *Mk3DT) ClearVoltageHistory(addr int) bool {
	this.execCmd(addr, "hstclear", "")
	// Nothing to return.
	return true
}

// todo
func (this *Mk3DT) TriggerLights(addr int) bool {
	this.execCmd(addr, "lights", "")
	// Return value as bool.
	return true
}

func (this *Mk3DT) GetMaxVoltage(addr int) float32 {
	r := this.execCmd(addr, "maxvolts", "")
	// Return the value as float32.
	return voltToFloat32(r.Value)
}

func (this *Mk3DT) GetMinVoltage(addr int) float32 {
	r := this.execCmd(addr, "minvolts", "")
	// Return the value as float32.
	return voltToFloat32(r.Value)
}

func (this *Mk3DT) GetStopChargeUnderVoltage(addr int) bool {
	r := this.execCmd(addr, "phev", "")
	// Return the value as bool.
	return r.Value == "1"
}

func (this *Mk3DT) SetStopChargeUnderVoltage(addr int, stop bool) bool {
	v := "0"
	if stop {
		v = "1"
	}
	r := this.execCmd(addr, "phev", v)
	// Return the value as bool.
	return r.Value == "1"
}

func (this *Mk3DT) GetRealTimeVoltage(addr int) float32 {
	r := this.execCmd(addr, "querytot", "")
	// Return the value as float32.
	return voltToFloat32(r.Value)
}

func (this *Mk3DT) GetLowVoltage(addr int) float32 {
	r := this.execCmd(addr, "readslow", "")
	// Return the value as float32.
	return voltToFloat32(r.Value)
}

// volts 0.000-9.999
func (this *Mk3DT) SetMaxVoltage(addr int, volts float32) bool {
	if volts < 0.000 || volts > 9.999 {
		return false
	}
	r := this.execCmd(addr, "sethigh", strconv.FormatFloat(float64(volts), 'f', 3, 32))
	// Check that the returned value is the same as the sent volts.
	return voltToFloat32(r.Value) == volts
}

// volts 0.000-9.999
func (this *Mk3DT) SetMinVoltage(addr int, volts float32) bool {
	if volts < 0.000 || volts > 9.999 {
		return false
	}
	r := this.execCmd(addr, "setlow", strconv.FormatFloat(float64(volts), 'f', 3, 32))
	// Check that the returned value is the same as the sent volts.
	return voltToFloat32(r.Value) == volts
}

// volts 0.000-9.999
func (this *Mk3DT) SetOverVoltage(addr int, volts float32) bool {
	if volts < 0.000 || volts > 9.999 {
		return false
	}
	r := this.execCmd(addr, "setover", strconv.FormatFloat(float64(volts), 'f', 3, 32))
	// Check that the returned value is the same as the sent volts.
	return voltToFloat32(r.Value) == volts
}

func (this *Mk3DT) GetAddrTemp(addr int) int {
	r := this.execCmd(addr, "temperat", "")
	// Return the value as int.
	return tempToInt(r.Value)
}

// temp 32-181
func (this *Mk3DT) SetFanMaxTemp(addr int, temp int) bool {
	if temp < 32 || temp > 180 {
		return false
	}
	r := this.execCmd(addr, "temphot", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return tempToInt(r.Value) == temp
}

// temp 32-181
func (this *Mk3DT) SetStopDissipatingTemp(addr int, temp int) bool {
	if temp < 32 || temp > 180 {
		return false
	}
	r := this.execCmd(addr, "tempoff", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return tempToInt(r.Value) == temp
}

// temp 32-181
func (this *Mk3DT) SetFanLowTemp(addr int, temp int) bool {
	if temp < 32 || temp > 180 {
		return false
	}
	r := this.execCmd(addr, "tempwarm", strconv.Itoa(temp))
	// Check that the returned value is the same as the sent temp.
	return tempToInt(r.Value) == temp
}

func (this *Mk3DT) GetCellsTemp(addr int) int {
	r := this.execCmd(addr, "xtrntemp", "")
	// Return the value as int (or string)?
	return tempToInt(r.Value)
}
