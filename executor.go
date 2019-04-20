package main

import ()

type Executor struct {
	mk3DT    *Mk3DT
	Commands *ExecutorCommands
}

type ExecutorCommands struct {
	SetStopChargeTemp         int     `yaml:"SetStopChargeTemp"`
	GetStopChargeTemp         bool    `yaml:"GetStopChargeTemp"`
	DisableStopChargeTemp     bool    `yaml:"DisableStopChargeTemp"`
	DisableShunt              bool    `yaml:"DisableShunt"`
	EnableShunt               bool    `yaml:"EnableShunt"`
	ForceFan                  int     `yaml:"ForceFan"`
	GetForceFan				  bool    `yaml:"GetForceFan"`
	GetFirstPosition          bool    `yaml:"GetFirstPosition"`
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
	GetAddrTemp               bool    `yaml:"GetAddrTemp"`
	SetFanMaxTemp             int     `yaml:"SetFanMaxTemp"`
	SetStopDissipatingTemp    int     `yaml:"SetStopDissipatingTemp"`
	SetFanLowTemp             int     `yaml:"SetFanLowTemp"`
	GetCellsTemp              bool    `yaml:"GetCellsTemp"`
}

type ExecutorCommandsResult struct {
	SetStopChargeTemp         bool    `yaml:"SetStopChargeTemp"`
	GetStopChargeTemp         int     `yaml:"GetStopChargeTemp"`
	DisableStopChargeTemp     bool    `yaml:"DisableStopChargeTemp"`
	DisableShunt              bool    `yaml:"DisableShunt"`
	EnableShunt               bool    `yaml:"EnableShunt"`
	ForceFan                  bool    `yaml:"ForceFan"`
	GetForceFan				  int     `yaml:"GetForceFan"`
	GetFirstPosition          bool    `yaml:"GetFirstPosition"`
	GetHighVoltage            float32 `yaml:"GetHighVoltage"`
	ClearMaxVoltageHistory    bool    `yaml:"ClearMaxVoltageHistory"`
	ClearMinVoltageHistory    bool    `yaml:"ClearMinVoltageHistory"`
	ClearVoltageHistory       bool    `yaml:"ClearVoltageHistory"`
	TriggerLights             bool    `yaml:"TriggerLights"`
	GetMaxVoltage             float32 `yaml:"GetMaxVoltage"`
	GetMinVoltage             float32 `yaml:"GetMinVoltage"`
	GetStopChargeUnderVoltage bool    `yaml:"GetStopChargeUnderVoltage"`
	SetStopChargeUnderVoltage bool    `yaml:"SetStopChargeUnderVoltage"`
	GetRealTimeVoltage        float32 `yaml:"GetRealTimeVoltage"`
	GetLowVoltage             float32 `yaml:"GetLowVoltage"`
	SetMaxVoltage             bool    `yaml:"SetMaxVoltage"`
	SetMinVoltage             bool    `yaml:"SetMinVoltage"`
	SetOverVoltage            bool    `yaml:"SetOverVoltage"`
	GetAddrTemp               int     `yaml:"GetAddrTemp"`
	SetFanMaxTemp             bool    `yaml:"SetFanMaxTemp"`
	SetStopDissipatingTemp    bool    `yaml:"SetStopDissipatingTemp"`
	SetFanLowTemp             bool    `yaml:"SetFanLowTemp"`
	GetCellsTemp              int     `yaml:"GetCellsTemp"`
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
		r.SetStopChargeTemp = this.mk3DT.SetStopChargeTemp(addr, this.Commands.SetStopChargeTemp)
	}
	if this.Commands.GetStopChargeTemp {
		r.GetStopChargeTemp = this.mk3DT.GetStopChargeTemp(addr)
	}
	if this.Commands.DisableStopChargeTemp {
		r.DisableStopChargeTemp = this.mk3DT.DisableStopChargeTemp(addr)
	}
	if this.Commands.DisableShunt {
		r.DisableShunt = this.mk3DT.DisableShunt(addr)
	}
	if this.Commands.EnableShunt {
		r.EnableShunt = this.mk3DT.EnableShunt(addr)
	}
	if this.Commands.ForceFan >= 0 && this.Commands.ForceFan <= 8 {
		r.ForceFan = this.mk3DT.ForceFan(addr, this.Commands.ForceFan)
	}
	if this.Commands.GetForceFan {
		r.GetForceFan = this.mk3DT.GetForceFan(addr)
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
	// If -1 turn off. If 1 turn on.
	if this.Commands.SetStopChargeUnderVoltage == -1 || this.Commands.SetStopChargeUnderVoltage == 1 {
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
	if this.Commands.GetAddrTemp {
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
