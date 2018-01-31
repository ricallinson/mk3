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

		It("should return 3.971 from GetMaxVolage()", func() {
			AssertEqual(mk3DT.GetMaxVolage(1), float32(3.971))
		})
	})

	Report(t)
}
