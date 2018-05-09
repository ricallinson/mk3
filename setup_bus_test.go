package main

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestSetupBus(t *testing.T) {

	var mockPort *MockPort
	var mk3DT *Mk3DT

	BeforeEach(func() {
		mockPort = NewMockPort()
		mk3DT = NewMk3DT(mockPort)
	})

	AfterEach(func() {
		mk3DT.Close()
		mockPort = nil
	})

	Describe("checkBus()", func() {

		It("should return 'true' from empty map", func() {
			cards := map[int]int{}
			mockPort.LastSerialNum = 0
			AssertEqual(checkBus(mk3DT, cards), true)
		})

		It("should return 'true' as found is a match to the given map", func() {
			cards := map[int]int{}
			cards[3] = 12
			cards[1] = 4
			cards[2] = 8
			mockPort.LastSerialNum = 3
			AssertEqual(checkBus(mk3DT, cards), true)
		})

		It("should return 'false' as found is more than given map", func() {
			cards := map[int]int{}
			cards[3] = 12
			cards[1] = 4
			cards[2] = 8
			mockPort.LastSerialNum = 4
			AssertEqual(checkBus(mk3DT, cards), false)
		})

	})

	Report(t)
}
