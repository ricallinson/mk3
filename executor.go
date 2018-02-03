package main

import ()

type Executor struct {
	mk3DT    *Mk3DT
	Commands *ExecutorCommands
}

type ExecutorCommands struct {
	SetStopChargeTemp         int     `yaml:"SetStopChargeTemp"`
	GetStopTemp               bool    `yaml:"GetStopTemp"`
	DisableStopChargeTemp     bool    `yaml:"DisableStopChargeTemp"`
	ChangeAddr                int     `yaml:"ChangeAddr"`
	GetCommands               bool    `yaml:"GetCommands"`
	DisableShunt              bool    `yaml:"DisableShunt"`
	EnableShunt               bool    `yaml:"EnableShunt"`
	ForceFan                  int     `yaml:"ForceFan"`
	GetFirstPosition          bool    `yaml:"GetFirstPosition"`
	SetFirstPosition          int     `yaml:"SetFirstPosition"`
	GetHighVoltage            bool    `yaml:"GetHighVoltage"`
	ClearMaxVoltageHistory    bool    `yaml:"ClearMaxVoltageHistory"`
	ClearMinVoltageHistory    bool    `yaml:"ClearMinVoltageHistory"`
	ClearVoltageHistory       bool    `yaml:"ClearVoltageHistory"`
	TriggerLights             bool    `yaml:"TriggerLights"`
	GetMaxVoltage             bool    `yaml:"GetMaxVoltage"`
	GetMinVoltage             bool    `yaml:"GetMinVoltage"`
	GetStopChargeUnderVoltage bool    `yaml:"GetStopChargeUnderVoltage"`
	SetStopChargeUnderVoltage int     `yaml:"SetStopChargeUnderVoltage"`
	GetRealTimeVoltage        bool    `yaml:"GetRealTimeVoltage"`
	GetLowVoltage             bool    `yaml:"GetLowVoltage"`
	SetMaxVoltage             float32 `yaml:"SetMaxVoltage"`
	SetMinVoltage             float32 `yaml:"SetMinVoltage"`
	SetOverVoltage            float32 `yaml:"SetOverVoltage"`
	GetStatus                 bool    `yaml:"GetStatus"`
	GetAddrTemp               int     `yaml:"GetAddrTemp"`
	SetFanMaxTemp             int     `yaml:"SetFanMaxTemp"`
	SetStopDissipatingTemp    int     `yaml:"SetStopDissipatingTemp"`
	SetFanLowTemp             int     `yaml:"SetFanLowTemp"`
	GetCellsTemp              bool    `yaml:"GetCellsTemp"`
}

type ExecutorCommandsResult struct {
	SetStopChargeTemp         bool
	GetStopTemp               int
	DisableStopChargeTemp     bool
	ChangeAddr                bool
	GetCommands               Mk3DTCommands
	DisableShunt              bool
	EnableShunt               bool
	ForceFan                  bool
	GetFirstPosition          bool
	SetFirstPosition          bool
	GetHighVoltage            float32
	ClearMaxVoltageHistory    bool
	ClearMinVoltageHistory    bool
	ClearVoltageHistory       bool
	TriggerLights             Mk3DTLightsStatus
	GetMaxVoltage             float32
	GetMinVoltage             float32
	GetStopChargeUnderVoltage bool
	SetStopChargeUnderVoltage bool
	GetRealTimeVoltage        float32
	GetLowVoltage             float32
	SetMaxVoltage             bool
	SetMinVoltage             bool
	SetOverVoltage            bool
	GetStatus                 Mk3DTStatus
	GetAddrTemp               int
	SetFanMaxTemp             bool
	SetStopDissipatingTemp    bool
	SetFanLowTemp             bool
	GetCellsTemp              string
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

func (this *Executor) ExecuteCommands() []*ExecutorCommandsResult {
	r := []*ExecutorCommandsResult{}
	for addr := 0; addr < 255; addr++ {
		r = append(r, this.ExecuteCommandsAtAddr(addr))
	}
	return r
}

func (this *Executor) ExecuteCommandsAtAddr(addr int) *ExecutorCommandsResult {
	r := &ExecutorCommandsResult{}
	if this.Commands.SetStopChargeTemp > 0 {
		r.SetStopChargeTemp = this.mk3DT.SetStopChargeTemp(addr, this.Commands.SetStopChargeTemp)
	}
	if this.Commands.GetStopTemp {
		r.GetStopTemp = this.mk3DT.GetStopTemp(addr)
	}
	if this.Commands.DisableStopChargeTemp {
		r.DisableStopChargeTemp = this.mk3DT.DisableStopChargeTemp(addr)
	}
	if this.Commands.ChangeAddr > -1 {
		r.ChangeAddr = this.mk3DT.ChangeAddr(addr, this.Commands.ChangeAddr)
	}
	if this.Commands.GetCommands {
		r.GetCommands = this.mk3DT.GetCommands(addr)
	}
	if this.Commands.DisableShunt {
		r.DisableShunt = this.mk3DT.DisableShunt(addr)
	}
	if this.Commands.EnableShunt {
		r.EnableShunt = this.mk3DT.EnableShunt(addr)
	}
	if this.Commands.ForceFan > -1 {
		r.ForceFan = this.mk3DT.ForceFan(addr, this.Commands.ForceFan)
	}
	if this.Commands.GetFirstPosition {
		r.GetFirstPosition = this.mk3DT.GetFirstPosition(addr)
	}
	if this.Commands.SetFirstPosition > -1 {
		r.SetFirstPosition = this.mk3DT.SetFirstPosition(addr, this.Commands.SetFirstPosition)
	}
	if this.Commands.GetHighVoltage {
		r.GetHighVoltage = this.mk3DT.GetHighVoltage(addr)
	}
	if this.Commands.ClearMaxVoltageHistory {
		r.ClearMaxVoltageHistory = this.mk3DT.ClearMaxVoltageHistory(addr)
	}
	if this.Commands.ClearMinVoltageHistory {
		r.ClearMinVoltageHistory = this.mk3DT.ClearMinVoltageHistory(addr)
	}
	if this.Commands.ClearVoltageHistory {
		r.ClearVoltageHistory = this.mk3DT.ClearVoltageHistory(addr)
	}
	if this.Commands.TriggerLights {
		r.TriggerLights = this.mk3DT.TriggerLights(addr)
	}
	if this.Commands.GetMaxVoltage {
		r.GetMaxVoltage = this.mk3DT.GetMaxVoltage(addr)
	}
	if this.Commands.GetMinVoltage {
		r.GetMinVoltage = this.mk3DT.GetMinVoltage(addr)
	}
	if this.Commands.GetStopChargeUnderVoltage {
		r.GetStopChargeUnderVoltage = this.mk3DT.GetStopChargeUnderVoltage(addr)
	}
	if this.Commands.SetStopChargeUnderVoltage > -1 {
		v := false
		if this.Commands.SetStopChargeUnderVoltage == 1 {
			v = true
		}
		r.SetStopChargeUnderVoltage = this.mk3DT.SetStopChargeUnderVoltage(addr, v)
	}
	if this.Commands.GetRealTimeVoltage {
		r.GetRealTimeVoltage = this.mk3DT.GetRealTimeVoltage(addr)
	}
	if this.Commands.GetLowVoltage {
		r.GetLowVoltage = this.mk3DT.GetLowVoltage(addr)
	}
	if this.Commands.SetMaxVoltage > 0 {
		r.SetMaxVoltage = this.mk3DT.SetMaxVoltage(addr, this.Commands.SetMaxVoltage)
	}
	if this.Commands.SetMinVoltage > 0 {
		r.SetMinVoltage = this.mk3DT.SetMinVoltage(addr, this.Commands.SetMinVoltage)
	}
	if this.Commands.SetOverVoltage > 0 {
		r.SetOverVoltage = this.mk3DT.SetOverVoltage(addr, this.Commands.SetOverVoltage)
	}
	if this.Commands.GetStatus {
		r.GetStatus = this.mk3DT.GetStatus(addr)
	}
	if this.Commands.GetAddrTemp > -1 {
		r.GetAddrTemp = this.mk3DT.GetAddrTemp(addr)
	}
	if this.Commands.SetFanMaxTemp > 0 {
		r.SetFanMaxTemp = this.mk3DT.SetFanMaxTemp(addr, this.Commands.SetFanMaxTemp)
	}
	if this.Commands.SetStopDissipatingTemp > 0 {
		r.SetStopDissipatingTemp = this.mk3DT.SetStopDissipatingTemp(addr, this.Commands.SetStopDissipatingTemp)
	}
	if this.Commands.SetFanLowTemp > 0 {
		r.SetFanLowTemp = this.mk3DT.SetFanLowTemp(addr, this.Commands.SetFanLowTemp)
	}
	if this.Commands.GetCellsTemp {
		r.GetCellsTemp = this.mk3DT.GetCellsTemp(addr)
	}
	return r
}
