package main

import (
	"encoding/json"
	"time"
)

type LiveData struct {
	Timestamp int64
	Address   int
	Volts     float32
	Temp      int
	SerialNum int
	CellCount int
}

func CreateLiveData(b []byte) *LiveData {
	var ld LiveData
	json.Unmarshal(b, &ld)
	return &ld
}

func GetRealtimeValues(mk3DT *Mk3DT, addr int) *LiveData {
	mk3DT.ClearVoltageHistory(addr)
	ld := &LiveData{}
	ld.Timestamp = time.Now().Unix()
	ld.Address = addr
	ld.CellCount = mk3DT.GetNumCells(addr)
	ld.Volts = mk3DT.GetMaxVoltage(addr)
	ld.Temp = mk3DT.GetCellsTemp(addr)
	ld.SerialNum = mk3DT.GetSerialNum(addr)
	return ld
}
