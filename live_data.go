package main

import (
	"encoding/json"
	"time"
)

type LiveData struct {
	Timestamp int64
	Address   int
	SerialNum int
	CellCount int
	Volts     float32
	MaxVolts  float32
	MinVolts  float32
	Temp      int
}

func CreateLiveData(b []byte) *LiveData {
	var ld LiveData
	json.Unmarshal(b, &ld)
	return &ld
}

func GetRealtimeValues(mk3DT *Mk3DT, addr int) *LiveData {
	ld := &LiveData{}
	ld.Timestamp = time.Now().Unix()
	ld.Address = addr
	ld.CellCount = mk3DT.GetNumCells(addr)
	ld.Volts = mk3DT.GetRealTimeVoltage(addr) / float32(ld.CellCount)
	ld.MaxVolts = mk3DT.GetMaxVoltageDetected(addr)
	ld.MinVolts = mk3DT.GetMinVoltageDetected(addr)
	ld.Temp = mk3DT.GetCellsTemp(addr)
	ld.SerialNum = mk3DT.GetSerialNum(addr)
	return ld
}
