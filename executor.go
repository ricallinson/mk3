package main

import (
	"gopkg.in/yaml.v2"
	"log"
)

type Executor struct {
	mk3DT  *Mk3DT
	ExeCmd *ExecutorCommands
}

type ExecutorCommands struct {
	SetStopChargeTemp         int
	GetStopTemp               bool
	DisableStopChargeTemp     bool
	ChangeAddr                int
	GetCommands               bool
	DisableShunt              bool
	EnableShunt               bool
	ForceFan                  int
	GetFirstPosition          bool
	SetFirstPosition          bool
	GetHighVoltage            bool
	ClearMaxVoltageHistory    bool
	ClearMinVoltageHistory    bool
	ClearVoltageHistory       bool
	TriggerLights             bool
	GetMaxVoltage             bool
	GetMinVoltage             bool
	GetStopChargeUnderVoltage bool
	SetStopChargeUnderVoltage bool
	GetRealTimeVoltage        bool
	GetLowVoltage             bool
	SetMaxVoltage             float32
	SetMinVoltage             float32
	SetOverVoltage            float32
	GetStatus                 bool
	GetAddrTemp               bool
	SetFanMaxTemp             int
	SetStopDissipatingTemp    int
	SetFanLowTemp             int
	GetCellsTemp              bool
}

type ExecutorCommandValues struct {
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
	GetCellsTemp              int
}

func NewExecutor(mk3DT *Mk3DT, p string, addr int) *Executor {
	this := &Executor{
		mk3DT: mk3DT,
	}
	err := yaml.Unmarshal(readFileToByteArray(p), &this.ExeCmd)
	if err != nil {
		log.Fatalf("YAML Error: %v", err)
	}
	log.Println(this.ExeCmd)
	return this
}

func (this *Executor) executeToAll() {
	for addr := 0; addr <= 255; addr++ {
		this.executeToAddr(addr)
	}
}

func (this *Executor) executeToAddr(addr int) {
	ecv := &ExecutorCommandValues{}
	switch {
	case this.ExeCmd.SetStopChargeTemp > 0:
		ecv.SetStopChargeTemp = this.mk3DT.SetStopChargeTemp(addr, this.ExeCmd.SetStopChargeTemp)
	case this.ExeCmd.GetStopTemp:
		ecv.GetStopTemp = this.mk3DT.GetStopTemp(addr)
	case this.ExeCmd.DisableStopChargeTemp:
		ecv.DisableStopChargeTemp = this.mk3DT.DisableStopChargeTemp(addr)
	case this.ExeCmd.ChangeAddr > -1:
		ecv.ChangeAddr = this.mk3DT.ChangeAddr(addr, this.ExeCmd.ChangeAddr)
	case this.ExeCmd.GetCommands:
		ecv.GetCommands = this.mk3DT.GetCommands(addr)
	case this.ExeCmd.DisableShunt:
		ecv.DisableShunt = this.mk3DT.DisableShunt(addr)
	case this.ExeCmd.EnableShunt:
		ecv.EnableShunt = this.mk3DT.EnableShunt(addr)
	case this.ExeCmd.ForceFan > -1:
		ecv.ForceFan = this.mk3DT.ForceFan(addr, this.ExeCmd.ForceFan)
	case this.ExeCmd.GetFirstPosition:
		ecv.GetFirstPosition = this.mk3DT.GetFirstPosition(addr)
	case this.ExeCmd.SetFirstPosition:
		ecv.SetFirstPosition = this.mk3DT.SetFirstPosition(addr, this.ExeCmd.SetFirstPosition)
	case this.ExeCmd.GetHighVoltage:
		ecv.GetHighVoltage = this.mk3DT.GetHighVoltage(addr)
	case this.ExeCmd.ClearMaxVoltageHistory:
		ecv.ClearMaxVoltageHistory = this.mk3DT.ClearMaxVoltageHistory(addr)
	case this.ExeCmd.ClearMinVoltageHistory:
		ecv.ClearMinVoltageHistory = this.mk3DT.ClearMinVoltageHistory(addr)
	case this.ExeCmd.ClearVoltageHistory:
		ecv.ClearVoltageHistory = this.mk3DT.ClearVoltageHistory(addr)
	case this.ExeCmd.TriggerLights:
		ecv.TriggerLights = this.mk3DT.TriggerLights(addr)
	case this.ExeCmd.GetMaxVoltage:
		ecv.GetMaxVoltage = this.mk3DT.GetMaxVoltage(addr)
	case this.ExeCmd.GetMinVoltage:
		ecv.GetMinVoltage = this.mk3DT.GetMinVoltage(addr)
	case this.ExeCmd.GetStopChargeUnderVoltage:
		ecv.GetStopChargeUnderVoltage = this.mk3DT.GetStopChargeUnderVoltage(addr)
	case this.ExeCmd.SetStopChargeUnderVoltage:
		ecv.SetStopChargeUnderVoltage = this.mk3DT.SetStopChargeUnderVoltage(addr, this.ExeCmd.SetStopChargeUnderVoltage)
	case this.ExeCmd.GetRealTimeVoltage:
		ecv.GetRealTimeVoltage = this.mk3DT.GetRealTimeVoltage(addr)
	case this.ExeCmd.GetLowVoltage:
		ecv.GetLowVoltage = this.mk3DT.GetLowVoltage(addr)
	case this.ExeCmd.SetMaxVoltage > 0:
		ecv.SetMaxVoltage = this.mk3DT.SetMaxVoltage(addr, this.ExeCmd.SetMaxVoltage)
	case this.ExeCmd.SetMinVoltage > 0:
		ecv.SetMinVoltage = this.mk3DT.SetMinVoltage(addr, this.ExeCmd.SetMinVoltage)
	case this.ExeCmd.SetOverVoltage > 0:
		ecv.SetOverVoltage = this.mk3DT.SetOverVoltage(addr, this.ExeCmd.SetOverVoltage)
	case this.ExeCmd.GetStatus:
		ecv.GetStatus = this.mk3DT.GetStatus(addr)
	case this.ExeCmd.GetAddrTemp:
		ecv.GetAddrTemp = this.mk3DT.GetAddrTemp(addr)
	case this.ExeCmd.SetFanMaxTemp > 0:
		ecv.SetFanMaxTemp = this.mk3DT.SetFanMaxTemp(addr, this.ExeCmd.SetFanMaxTemp)
	case this.ExeCmd.SetStopDissipatingTemp > 0:
		ecv.SetStopDissipatingTemp = this.mk3DT.SetStopDissipatingTemp(addr, this.ExeCmd.SetStopDissipatingTemp)
	case this.ExeCmd.SetFanLowTemp > 0:
		ecv.SetFanLowTemp = this.mk3DT.SetFanLowTemp(addr, this.ExeCmd.SetFanLowTemp)
	case this.ExeCmd.GetCellsTemp:
		ecv.GetCellsTemp = this.mk3DT.GetCellsTemp(addr)
	}
}
