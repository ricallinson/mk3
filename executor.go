package main

import (
	"gopkg.in/yaml.v2"
	"log"
)

type Executor struct {
	mk3DT  *Mk3DT
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
	SetFirstPosition          bool    `yaml:"SetFirstPosition"`
	GetHighVoltage            bool    `yaml:"GetHighVoltage"`
	ClearMaxVoltageHistory    bool    `yaml:"ClearMaxVoltageHistory"`
	ClearMinVoltageHistory    bool    `yaml:"ClearMinVoltageHistory"`
	ClearVoltageHistory       bool    `yaml:"ClearVoltageHistory"`
	TriggerLights             bool    `yaml:"TriggerLights"`
	GetMaxVoltage             bool    `yaml:"GetMaxVoltage"`
	GetMinVoltage             bool    `yaml:"GetMinVoltage"`
	GetStopChargeUnderVoltage bool    `yaml:"GetStopChargeUnderVoltage"`
	SetStopChargeUnderVoltage bool    `yaml:"SetStopChargeUnderVoltage"`
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

func NewExecutor(mk3DT *Mk3DT, p string) *Executor {
	this := &Executor{
		mk3DT:  mk3DT,
		Commands: &ExecutorCommands{},
	}
	if p != "" {
		err := yaml.Unmarshal(readFileToByteArray(p), &this.Commands)
		if err != nil {
			log.Fatalf("YAML Error: %v", err)
		}
	}
	return this
}

func (this *Executor) Run() {
	for addr := 0; addr <= 0; addr++ {
		this.executeToAddr(addr)
	}
}

func (this *Executor) executeToAddr(addr int) {
	ecv := &ExecutorCommandValues{}

	log.Println(this.Commands.GetStopTemp)

	if this.Commands.SetStopChargeTemp > 0 {
		ecv.SetStopChargeTemp = this.mk3DT.SetStopChargeTemp(addr, this.Commands.SetStopChargeTemp)
	}
	if this.Commands.GetStopTemp {
		ecv.GetStopTemp = this.mk3DT.GetStopTemp(addr)
	}
	if this.Commands.DisableStopChargeTemp {
		ecv.DisableStopChargeTemp = this.mk3DT.DisableStopChargeTemp(addr)
	}
	if this.Commands.ChangeAddr > -1 {
		ecv.ChangeAddr = this.mk3DT.ChangeAddr(addr, this.Commands.ChangeAddr)
	}
	if this.Commands.GetCommands {
		ecv.GetCommands = this.mk3DT.GetCommands(addr)
	}
	if this.Commands.DisableShunt {
		ecv.DisableShunt = this.mk3DT.DisableShunt(addr)
	}
	if this.Commands.EnableShunt {
		ecv.EnableShunt = this.mk3DT.EnableShunt(addr)
	}
	if this.Commands.ForceFan > -1 {
		ecv.ForceFan = this.mk3DT.ForceFan(addr, this.Commands.ForceFan)
	}
	if this.Commands.GetFirstPosition {
		ecv.GetFirstPosition = this.mk3DT.GetFirstPosition(addr)
	}
	if this.Commands.SetFirstPosition {
		ecv.SetFirstPosition = this.mk3DT.SetFirstPosition(addr, this.Commands.SetFirstPosition)
	}
	if this.Commands.GetHighVoltage {
		ecv.GetHighVoltage = this.mk3DT.GetHighVoltage(addr)
	}
	if this.Commands.ClearMaxVoltageHistory {
		ecv.ClearMaxVoltageHistory = this.mk3DT.ClearMaxVoltageHistory(addr)
	}
	if this.Commands.ClearMinVoltageHistory {
		ecv.ClearMinVoltageHistory = this.mk3DT.ClearMinVoltageHistory(addr)
	}
	if this.Commands.ClearVoltageHistory {
		ecv.ClearVoltageHistory = this.mk3DT.ClearVoltageHistory(addr)
	}
	if this.Commands.TriggerLights {
		ecv.TriggerLights = this.mk3DT.TriggerLights(addr)
	}
	if this.Commands.GetMaxVoltage {
		ecv.GetMaxVoltage = this.mk3DT.GetMaxVoltage(addr)
	}
	if this.Commands.GetMinVoltage {
		ecv.GetMinVoltage = this.mk3DT.GetMinVoltage(addr)
	}
	if this.Commands.GetStopChargeUnderVoltage {
		ecv.GetStopChargeUnderVoltage = this.mk3DT.GetStopChargeUnderVoltage(addr)
	}
	if this.Commands.SetStopChargeUnderVoltage {
		ecv.SetStopChargeUnderVoltage = this.mk3DT.SetStopChargeUnderVoltage(addr, this.Commands.SetStopChargeUnderVoltage)
	}
	if this.Commands.GetRealTimeVoltage {
		ecv.GetRealTimeVoltage = this.mk3DT.GetRealTimeVoltage(addr)
	}
	if this.Commands.GetLowVoltage {
		ecv.GetLowVoltage = this.mk3DT.GetLowVoltage(addr)
	}
	if this.Commands.SetMaxVoltage > 0 {
		ecv.SetMaxVoltage = this.mk3DT.SetMaxVoltage(addr, this.Commands.SetMaxVoltage)
	}
	if this.Commands.SetMinVoltage > 0 {
		ecv.SetMinVoltage = this.mk3DT.SetMinVoltage(addr, this.Commands.SetMinVoltage)
	}
	if this.Commands.SetOverVoltage > 0 {
		ecv.SetOverVoltage = this.mk3DT.SetOverVoltage(addr, this.Commands.SetOverVoltage)
	}
	if this.Commands.GetStatus {
		ecv.GetStatus = this.mk3DT.GetStatus(addr)
	}
	if this.Commands.GetAddrTemp > -1 {
		ecv.GetAddrTemp = this.mk3DT.GetAddrTemp(addr)
	}
	if this.Commands.SetFanMaxTemp > 0 {
		ecv.SetFanMaxTemp = this.mk3DT.SetFanMaxTemp(addr, this.Commands.SetFanMaxTemp)
	}
	if this.Commands.SetStopDissipatingTemp > 0 {
		ecv.SetStopDissipatingTemp = this.mk3DT.SetStopDissipatingTemp(addr, this.Commands.SetStopDissipatingTemp)
	}
	if this.Commands.SetFanLowTemp > 0 {
		ecv.SetFanLowTemp = this.mk3DT.SetFanLowTemp(addr, this.Commands.SetFanLowTemp)
	}
	if this.Commands.GetCellsTemp {
		ecv.GetCellsTemp = this.mk3DT.GetCellsTemp(addr)
	}
	// b, _ := yaml.Marshal(ecv)
	// log.Println(string(b))
}
