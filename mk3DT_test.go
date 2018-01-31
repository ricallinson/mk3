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
			AssertEqual(mk3DT.SetStopTemp(1, 189), true)
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

		It("should return 3.971 from GetMaxVolage()", func() {
			AssertEqual(mk3DT.GetMaxVolage(1), float32(3.971))
		})

		It("should return 2.432 from GetMinVolage()", func() {
			AssertEqual(mk3DT.GetMinVolage(1), float32(2.432))
		})

		It("should return 'true' from GetChargeUnderVoltage()", func() {
			AssertEqual(mk3DT.GetStopChargeUnderVoltage(1), false)
		})

		It("should return the sent value from SetChargeUnderVoltage()", func() {
			AssertEqual(mk3DT.SetStopChargeUnderVoltage(1, false), false)
			AssertEqual(mk3DT.SetStopChargeUnderVoltage(1, true), true)
		})
	})

	Report(t)
}
