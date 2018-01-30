package main

import (
	"strconv"
	"bytes"
	"log"
)

type Mk3 struct {
	serialPort SerialPort
}

func NewMk3(p SerialPort) *Mk3 {
	this := &Mk3{
		serialPort: p,
	}
	return this
}

func (this *Mk3) readBytes(delim byte) []byte {
	limit := 10000
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

func (this *Mk3) execCmd(unit int, cmd string, value string) string {
	this.readBytes(0) // Clear the buffer.
	this.writeBytes([]byte(strconv.Itoa(unit) + cmd + "." + value + "\n\r"))
	return string(this.readBytes(0))
}

// temp 32-180 F
func (this *Mk3) SetStopTemp(unit int, temp int) {
	this.execCmd(unit, "bt", strconv.Itoa(temp))
}

func (this *Mk3) DisableStopTemp(unit int) {
	this.execCmd(unit, "btd", "")
}

// addr 0-255
func (this *Mk3) ChangeAddr(unit int, addr int) {
	this.execCmd(unit, "ch", strconv.Itoa(addr))
}

func (this *Mk3) GetCommands(unit int) {
	this.execCmd(unit, "", "")
}

func (this *Mk3) DisableShunt(unit int) {
	this.execCmd(unit, "d", "")
}

func (this *Mk3) EnableShunt(unit int) {
	this.execCmd(unit, "e", "")
}

// level 0-8
func (this *Mk3) ForceFan(unit int, level int) {
	this.execCmd(unit, "f", strconv.Itoa(level))
}

func (this *Mk3) GetFirstPosition(unit int) {
	this.execCmd(unit, "fi", "")
}

func (this *Mk3) SetFirstPosition(unit int, positive bool) {
	this.execCmd(unit, "fi", strconv.FormatBool(positive))
}

func (this *Mk3) GetHighVoltage(unit int) {
	this.execCmd(unit, "g", "")
}

func (this *Mk3) ClearMaxVolageHistory(unit int) {
	this.execCmd(unit, "hma", "")
}

func (this *Mk3) ClearMinVolageHistory(unit int) {
	this.execCmd(unit, "hmi", "")
}

func (this *Mk3) ClearVolageHistory(unit int) {
	this.execCmd(unit, "h", "")
}

func (this *Mk3) TriggerLights(unit int) {
	this.execCmd(unit, "l", "")
}

func (this *Mk3) GetMaxVolage(unit int) {
	this.execCmd(unit, "ma", "")
}

func (this *Mk3) GetMinVolage(unit int) {
	this.execCmd(unit, "mi", "")
}

func (this *Mk3) SetStopChargeUnderVoltage(unit int, stop bool) {
	this.execCmd(unit, "p", strconv.FormatBool(stop))
}

func (this *Mk3) GetRealTimeVoltage(unit int) {
	this.execCmd(unit, "q", "")
}

func (this *Mk3) GetLowVoltage(unit int) {
	this.execCmd(unit, "r", "")
}

// volts 0.000-9.999
func (this *Mk3) SetMaxVoltage(unit int, volts int) {
	this.execCmd(unit, "seth", strconv.Itoa(volts))
}

// volts 0.000-9.999
func (this *Mk3) SetMinVoltage(unit int, volts int) {
	this.execCmd(unit, "setl", strconv.Itoa(volts))
}

// volts 0.000-9.999
func (this *Mk3) SetOverVoltage(unit int, volts int) {
	this.execCmd(unit, "seto", strconv.Itoa(volts))
}

func (this *Mk3) GetStatus(unit int) {
	this.execCmd(unit, "s", "")
}

func (this *Mk3) GetUintTemp(unit int) {
	this.execCmd(unit, "t", "")
}

// temp 32-181
func (this *Mk3) SetFanMaxTemp(unit int, temp int) {
	this.execCmd(unit, "temph", strconv.Itoa(temp))
}

// temp 32-181
func (this *Mk3) SetStopDissipatingTemp(unit int, temp int) {
	this.execCmd(unit, "tempo", strconv.Itoa(temp))
}

// temp 32-181
func (this *Mk3) SetFanLowTemp(unit int, temp int) {
	this.execCmd(unit, "tempw", strconv.Itoa(temp))
}

func (this *Mk3) GetCellsTemp(unit int) {
	this.execCmd(unit, "x", "")
}
