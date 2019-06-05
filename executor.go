package main

import (
	"fmt"
)

type Executor struct {
	mk3DT    *Mk3DT
	Commands *ExecutorCommands
}

type ExecutorCommands struct {
	SetStopChargeTemp            int     `yaml:"SetStopChargeTemp"`
	GetStopChargeTemp            bool    `yaml:"GetStopChargeTemp"`
	GetSerialNum                 bool    `yaml:"GetSerialNum"`
	GetNumCells                  bool    `yaml:"GetNumCells"`
	GetCellRange                 bool    `yaml:"GetCellRange"`
	DisableStopChargeTemp        bool    `yaml:"DisableStopChargeTemp"`
	GetShuntMode                 bool    `yaml:"GetShuntMode"`
	DisableShunt                 bool    `yaml:"DisableShunt"`
	EnableShunt                  bool    `yaml:"EnableShunt"`
	SetForceFan                  int     `yaml:"SetForceFan"`
	GetForceFan                  bool    `yaml:"GetForceFan"`
	GetFirstPosition             bool    `yaml:"GetFirstPosition"`
	GetHighVoltage               bool    `yaml:"GetHighVoltage"`
	ClearMaxVoltageHistory       bool    `yaml:"ClearMaxVoltageHistory"`
	ClearMinVoltageHistory       bool    `yaml:"ClearMinVoltageHistory"`
	ClearVoltageHistory          bool    `yaml:"ClearVoltageHistory"`
	TriggerLights                bool    `yaml:"TriggerLights"`
	GetMaxVoltageDetected        bool    `yaml:"GetMaxVoltageDetected"`
	GetMinVoltageDetected        bool    `yaml:"GetMinVoltageDetected"`
	GetStopChargeUnderVoltage    bool    `yaml:"GetStopChargeUnderVoltage"`
	SetStopChargeUnderVoltageOn  bool    `yaml:"SetStopChargeUnderVoltageOn"`
	SetStopChargeUnderVoltageOff bool    `yaml:"SetStopChargeUnderVoltageOff"`
	GetRealTimeVoltage           bool    `yaml:"GetRealTimeVoltage"`
	GetLowVoltage                bool    `yaml:"GetLowVoltage"`
	SetMaxVoltage                float32 `yaml:"SetMaxVoltage"`
	SetMinVoltage                float32 `yaml:"SetMinVoltage"`
	SetOverVoltage               float32 `yaml:"SetOverVoltage"`
	GetAddrTemp                  bool    `yaml:"GetAddrTemp"`
	SetFanMaxTemp                int     `yaml:"SetFanMaxTemp"`
	SetStopDissipatingTemp       int     `yaml:"SetStopDissipatingTemp"`
	SetFanLowTemp                int     `yaml:"SetFanLowTemp"`
	GetCellsTemp                 bool    `yaml:"GetCellsTemp"`
}

type ExecutorCommandsResult struct {
	StopChargeTemp         int      `yaml:"StopChargeTemp"`
	SerialNum              int      `yaml:"SerialNum"`
	NumCells               int      `yaml:"NumCells"`
	CellRange              string   `yaml:"CellRange"`
	DisableStopChargeTemp  bool     `yaml:"DisableStopChargeTemp"`
	ShuntMode              bool     `yaml:"ShuntMode"`
	ForceFan               int      `yaml:"ForceFan"`
	FirstPosition          bool     `yaml:"FirstPosition"`
	HighVoltage            float32  `yaml:"HighVoltage"`
	MaxVoltageDetected     float32  `yaml:"MaxVoltageDetected"`
	MinVoltageDetected     float32  `yaml:"MinVoltageDetected"`
	StopChargeUnderVoltage bool     `yaml:"StopChargeUnderVoltage"`
	RealTimeVoltage        float32  `yaml:"RealTimeVoltage"`
	LowVoltage             float32  `yaml:"LowVoltage"`
	MaxVoltage             float32  `yaml:"MaxVoltage"`
	MinVoltage             float32  `yaml:"MinVoltage"`
	OverVoltage            float32  `yaml:"OverVoltage"`
	AddrTemp               int      `yaml:"AddrTemp"`
	FanMaxTemp             int      `yaml:"FanMaxTemp"`
	StopDissipatingTemp    int      `yaml:"StopDissipatingTemp"`
	FanLowTemp             int      `yaml:"FanLowTemp"`
	CellsTemp              int      `yaml:"CellsTemp"`
	Commands               []string `yaml:"Commands"`
}

func NewExecutor(mk3DT *Mk3DT) *Executor {
	this := &Executor{
		mk3DT:    mk3DT,
		Commands: &ExecutorCommands{},
	}
	return this
}

func (this *Executor) Close() {
	this.mk3DT.Close()
}

func (this *Executor) ExecuteCommands(to int) map[int]*ExecutorCommandsResult {
	r := map[int]*ExecutorCommandsResult{}
	for addr := 0; addr <= to; addr++ {
		if this.mk3DT.GetStopChargeTemp(addr) > 0 {
			r[addr] = this.ExecuteCommandsAtAddr(addr)
		}
	}
	return r
}

func (this *Executor) ExecuteCommandsAtAddr(addr int) *ExecutorCommandsResult {
	r := &ExecutorCommandsResult{}
	if this.Commands.SetStopChargeTemp > 0 {
		if this.mk3DT.SetStopChargeTemp(addr, this.Commands.SetStopChargeTemp) {
			r.Commands = append(r.Commands, "SetStopChargeTemp")
		}
	}
	if this.Commands.GetStopChargeTemp {
		r.StopChargeTemp = this.mk3DT.GetStopChargeTemp(addr)
		r.Commands = append(r.Commands, "GetStopChargeTemp")
	}
	if this.Commands.GetSerialNum {
		r.SerialNum = this.mk3DT.GetSerialNum(addr)
		r.Commands = append(r.Commands, "GetSerialNum")
	}
	if this.Commands.GetNumCells {
		r.NumCells = this.mk3DT.GetNumCells(addr)
		r.Commands = append(r.Commands, "GetNumCells")
	}
	if this.Commands.GetCellRange {
		start, end := this.mk3DT.GetCellRange(addr)
		r.CellRange = fmt.Sprintf("%d %d", start, end)
		r.Commands = append(r.Commands, "GetCellRange")
	}
	if this.Commands.DisableStopChargeTemp {
		if this.mk3DT.DisableStopChargeTemp(addr) {
			r.Commands = append(r.Commands, "DisableStopChargeTemp")
		}
	}
	if this.Commands.GetShuntMode {
		r.ShuntMode = this.mk3DT.GetShuntMode(addr)
		r.Commands = append(r.Commands, "GetShuntMode")
	}
	if this.Commands.DisableShunt {
		if this.mk3DT.DisableShunt(addr) {
			r.Commands = append(r.Commands, "DisableShunt")
		}
	}
	if this.Commands.EnableShunt {
		if this.mk3DT.EnableShunt(addr) {
			r.Commands = append(r.Commands, "EnableShunt")
		}
	}
	if this.Commands.SetForceFan >= 0 && this.Commands.SetForceFan <= 8 {
		if this.mk3DT.ForceFan(addr, this.Commands.SetForceFan) {
			r.Commands = append(r.Commands, "SetForceFan")
		}
	}
	if this.Commands.GetForceFan {
		r.ForceFan = this.mk3DT.GetForceFan(addr)
		r.Commands = append(r.Commands, "GetForceFan")
	}
	if this.Commands.GetHighVoltage {
		r.HighVoltage = this.mk3DT.GetHighVoltage(addr)
		r.Commands = append(r.Commands, "GetHighVoltage")
	}
	if this.Commands.ClearMaxVoltageHistory {
		if this.mk3DT.ClearMaxVoltageHistory(addr) {
			r.Commands = append(r.Commands, "ClearMaxVoltageHistory")
		}
	}
	if this.Commands.ClearMinVoltageHistory {
		if this.mk3DT.ClearMinVoltageHistory(addr) {
			r.Commands = append(r.Commands, "ClearMinVoltageHistory")
		}
	}
	if this.Commands.ClearVoltageHistory {
		if this.mk3DT.ClearVoltageHistory(addr) {
			r.Commands = append(r.Commands, "ClearVoltageHistory")
		}
	}
	if this.Commands.TriggerLights {
		if this.mk3DT.TriggerLights(addr) {
			r.Commands = append(r.Commands, "TriggerLights")
		}
	}
	if this.Commands.GetMaxVoltageDetected {
		r.MaxVoltageDetected = this.mk3DT.GetMaxVoltageDetected(addr)
		r.Commands = append(r.Commands, "GetMaxVoltageDetected")
	}
	if this.Commands.GetMinVoltageDetected {
		r.MinVoltageDetected = this.mk3DT.GetMinVoltageDetected(addr)
		r.Commands = append(r.Commands, "GetMinVoltageDetected")
	}
	if this.Commands.GetStopChargeUnderVoltage {
		r.StopChargeUnderVoltage = this.mk3DT.GetStopChargeUnderVoltage(addr)
		r.Commands = append(r.Commands, "GetStopChargeUnderVoltage")
	}
	if this.Commands.SetStopChargeUnderVoltageOn {
		if this.mk3DT.SetStopChargeUnderVoltageOn(addr) {
			r.Commands = append(r.Commands, "SetStopChargeUnderVoltageOn")
		}
	}
	if this.Commands.SetStopChargeUnderVoltageOff {
		if this.mk3DT.SetStopChargeUnderVoltageOff(addr) {
			r.Commands = append(r.Commands, "SetStopChargeUnderVoltageOff")
		}
	}
	if this.Commands.GetRealTimeVoltage {
		r.RealTimeVoltage = this.mk3DT.GetRealTimeVoltage(addr)
		r.Commands = append(r.Commands, "GetRealTimeVoltage")
	}
	if this.Commands.GetLowVoltage {
		r.LowVoltage = this.mk3DT.GetLowVoltage(addr)
		r.Commands = append(r.Commands, "GetLowVoltage")
	}
	if this.Commands.SetMaxVoltage > 0 {
		if this.mk3DT.SetMaxVoltage(addr, this.Commands.SetMaxVoltage) {
			r.Commands = append(r.Commands, "SetMaxVoltage")
		}
	}
	if this.Commands.SetMinVoltage > 0 {
		if this.mk3DT.SetMinVoltage(addr, this.Commands.SetMinVoltage) {
			r.Commands = append(r.Commands, "SetMinVoltage")
		}
	}
	if this.Commands.SetOverVoltage > 0 {
		if this.mk3DT.SetOverVoltage(addr, this.Commands.SetOverVoltage) {
			r.Commands = append(r.Commands, "SetOverVoltage")
		}
	}
	if this.Commands.GetAddrTemp {
		r.AddrTemp = this.mk3DT.GetAddrTemp(addr)
		r.Commands = append(r.Commands, "GetAddrTemp")
	}
	if this.Commands.SetFanMaxTemp > 0 {
		if this.mk3DT.SetFanMaxTemp(addr, this.Commands.SetFanMaxTemp) {
			r.Commands = append(r.Commands, "SetFanMaxTemp")
		}
	}
	if this.Commands.SetStopDissipatingTemp > 0 {
		if this.mk3DT.SetStopDissipatingTemp(addr, this.Commands.SetStopDissipatingTemp) {
			r.Commands = append(r.Commands, "SetStopDissipatingTemp")
		}
	}
	if this.Commands.SetFanLowTemp > 0 {
		if this.mk3DT.SetFanLowTemp(addr, this.Commands.SetFanLowTemp) {
			r.Commands = append(r.Commands, "SetFanLowTemp")
		}
	}
	if this.Commands.GetCellsTemp {
		r.CellsTemp = this.mk3DT.GetCellsTemp(addr)
		r.Commands = append(r.Commands, "GetCellsTemp")
	}
	return r
}
