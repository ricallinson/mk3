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
	ClearMaxVolageHistory     bool
	ClearMinVolageHistory     bool
	ClearVolageHistory        bool
	TriggerLights             Mk3DTLightsStatus
	GetMaxVolage              float32
	GetMinVolage              float32
	GetStopChargeUnderVoltage bool
	SetStopChargeUnderVoltage bool
	GetRealTimeVoltage        float32
	GetLowVoltage             float32
	SetMaxVoltage             bool
	SetMinVoltage             bool
	SetOverVoltage            bool
	GetStatus                 Mk3DTStatus
	GetAddrTemp               bool
	SetFanMaxTemp             bool
	SetStopDissipatingTemp    bool
	SetFanLowTemp             bool
	GetCellsTemp              int
}

func NewExecutor(mk3DT *Mk3DT, p string) *Executor {
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

func (this *Executor) writeAll() {

}

func (this *Executor) writeAddr() {

}

func (this *Executor) read() map[int]*ExecutorCommands {
	return map[int]*ExecutorCommands{}
}
