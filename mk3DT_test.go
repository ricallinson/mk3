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

		It("should return the sent value from DisableStopTemp()", func() {
			AssertEqual(mk3DT.DisableStopTemp(1), true)
		})

		It("should return the sent value from ChangeAddr()", func() {
			AssertEqual(mk3DT.ChangeAddr(1, 100), true)
		})

		// Here

		It("should return 3.971 from GetMaxVolage()", func() {
			AssertEqual(mk3DT.GetMaxVolage(1), float32(3.971))
		})
	})

	Report(t)
}
