package main

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestMk3DT(t *testing.T) {

	var mk3DT *Mk3DT

	BeforeEach(func() {
		mk3DT = NewMk3DT(NewMockPort())
	})

	Describe("Mk3DT()", func() {

		It("should return a Mk3DT object", func() {
			AssertEqual(reflect.TypeOf(mk3DT).String(), "*main.Mk3DT")
		})

		It("should return the sent value from SetStopTemp()", func() {
			AssertEqual(mk3DT.SetStopTemp(1, 31), false)
			AssertEqual(mk3DT.SetStopTemp(1, 181), false)
			AssertEqual(mk3DT.SetStopTemp(1, 170), true)
		})

		It("should return the sent value of GetStopTemp()", func() {
			AssertEqual(mk3DT.GetStopTemp(1), 180)
		})

		It("should return 'true' from DisableStopTemp()", func() {
			AssertEqual(mk3DT.DisableStopTemp(1), true)
		})

		It("should return the sent value from ChangeAddr()", func() {
			AssertEqual(mk3DT.ChangeAddr(1, 100), true)
		})

		// Todo
		It("should return the Commands from GetCommands()", func() {
			AssertEqual(mk3DT.GetCommands(1), true)
		})

		It("should return 'true' from DisableShunt()", func() {
			AssertEqual(mk3DT.DisableShunt(1), true)
		})

		It("should return 'true' from EnableShunt()", func() {
			AssertEqual(mk3DT.EnableShunt(1), true)
		})

		It("should return 'true' from ForceFan() if it's in range", func() {
			AssertEqual(mk3DT.ForceFan(1, -9), false)
			AssertEqual(mk3DT.ForceFan(1, 9), false)
			AssertEqual(mk3DT.ForceFan(1, 5), true)
		})

		It("should return 'true' from GetFirstPosition()", func() {
			AssertEqual(mk3DT.GetFirstPosition(1), true)
		})

		It("should return 'true' from SetFirstPosition()", func() {
			AssertEqual(mk3DT.SetFirstPosition(1, false), true)
			AssertEqual(mk3DT.SetFirstPosition(1, true), true)
		})

		It("should return 'true' from GetHighVoltage()", func() {
			AssertEqual(mk3DT.GetHighVoltage(1), float32(3.9))
		})

		It("should return 'true' from ClearMaxVolageHistory()", func() {
			AssertEqual(mk3DT.ClearMaxVolageHistory(1), true)
		})

		It("should return 'true' from ClearMinVolageHistory()", func() {
			AssertEqual(mk3DT.ClearMinVolageHistory(1), true)
		})

		It("should return 'true' from ClearVolageHistory()", func() {
			AssertEqual(mk3DT.ClearVolageHistory(1), true)
		})

		// Todo
		It("should return the LightsStatus from TriggerLights()", func() {
			AssertEqual(mk3DT.TriggerLights(1), true)
		})

		It("should return 3.971 from GetMaxVoltage()", func() {
			AssertEqual(mk3DT.GetMaxVoltage(1), float32(3.971))
		})

		It("should return 2.432 from GetMinVoltage()", func() {
			AssertEqual(mk3DT.GetMinVoltage(1), float32(2.432))
		})

		It("should return 'true' from GetChargeUnderVoltage()", func() {
			AssertEqual(mk3DT.GetStopChargeUnderVoltage(1), false)
		})

		It("should return the sent value from SetChargeUnderVoltage()", func() {
			AssertEqual(mk3DT.SetStopChargeUnderVoltage(1, false), false)
			AssertEqual(mk3DT.SetStopChargeUnderVoltage(1, true), true)
		})

		It("should return 2.432 from GetLowVoltage()", func() {
			AssertEqual(mk3DT.GetLowVoltage(1), float32(2.432))
		})

		It("should return 'true' from SetMaxVoltage()", func() {
			AssertEqual(mk3DT.SetMaxVoltage(1, 0.0001), false)
			AssertEqual(mk3DT.SetMaxVoltage(1, 9.9991), false)
			AssertEqual(mk3DT.SetMaxVoltage(1, 3.123), true)
		})

		It("should return 'true' from SetMinVoltage()", func() {
			AssertEqual(mk3DT.SetMinVoltage(1, 0.0001), false)
			AssertEqual(mk3DT.SetMinVoltage(1, 9.9991), false)
			AssertEqual(mk3DT.SetMinVoltage(1, 2.123), true)
		})

		It("should return 'true' from SetOverVoltage()", func() {
			AssertEqual(mk3DT.SetOverVoltage(1, 0.0001), false)
			AssertEqual(mk3DT.SetOverVoltage(1, 9.9991), false)
			AssertEqual(mk3DT.SetOverVoltage(1, 3.321), true)
		})

		// Todo
		It("should return the Status from GetStatus()", func() {
			AssertEqual(mk3DT.GetStatus(1), true)
		})

		It("should return 120 from GetAddrTemp()", func() {
			AssertEqual(mk3DT.GetAddrTemp(1), 120)
		})

		It("should return the sent value from SetFanMaxTemp()", func() {
			AssertEqual(mk3DT.SetFanMaxTemp(1, 31), false)
			AssertEqual(mk3DT.SetFanMaxTemp(1, 181), false)
			AssertEqual(mk3DT.SetFanMaxTemp(1, 170), true)
		})

		It("should return the sent value from SetStopDissipatingTemp()", func() {
			AssertEqual(mk3DT.SetStopDissipatingTemp(1, 31), false)
			AssertEqual(mk3DT.SetStopDissipatingTemp(1, 181), false)
			AssertEqual(mk3DT.SetStopDissipatingTemp(1, 170), true)
		})

		It("should return the sent value from SetFanLowTemp()", func() {
			AssertEqual(mk3DT.SetFanLowTemp(1, 31), false)
			AssertEqual(mk3DT.SetFanLowTemp(1, 181), false)
			AssertEqual(mk3DT.SetFanLowTemp(1, 170), true)
		})

		It("should return 0 from GetCellsTemp()", func() {
			AssertEqual(mk3DT.GetCellsTemp(1), 0)
		})
	})

	Report(t)
}
