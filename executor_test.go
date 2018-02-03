package main

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestExecutor(t *testing.T) {

	Describe("Executor()", func() {

		It("should return an Executor object", func() {
			e := NewExecutor(NewMk3DT(NewMockPort()), "")
			AssertEqual(reflect.TypeOf(e).String(), "*main.Executor")
		})

		It("should populate the ExecutorCommands object from YAML", func() {
			e := NewExecutor(NewMk3DT(NewMockPort()), "./fixtures/all_commands.yaml")
			AssertEqual(e.Commands.SetStopChargeTemp, 120)
			AssertEqual(e.Commands.GetStopTemp, true)
			AssertEqual(e.Commands.DisableStopChargeTemp, true)
			AssertEqual(e.Commands.ChangeAddr, 1)
			AssertEqual(e.Commands.GetCommands, true)
			AssertEqual(e.Commands.DisableShunt, true)
			AssertEqual(e.Commands.EnableShunt, true)
			AssertEqual(e.Commands.ForceFan, 4)
			AssertEqual(e.Commands.GetFirstPosition, true)
			AssertEqual(e.Commands.SetFirstPosition, false)
			AssertEqual(e.Commands.GetHighVoltage, true)
			AssertEqual(e.Commands.ClearMaxVoltageHistory, true)
			AssertEqual(e.Commands.ClearMinVoltageHistory, true)
			AssertEqual(e.Commands.ClearVoltageHistory, true)
			AssertEqual(e.Commands.TriggerLights, true)
			AssertEqual(e.Commands.GetMaxVoltage, true)
			AssertEqual(e.Commands.GetMinVoltage, true)
			AssertEqual(e.Commands.GetStopChargeUnderVoltage, true)
			AssertEqual(e.Commands.SetStopChargeUnderVoltage, true)
			AssertEqual(e.Commands.GetRealTimeVoltage, true)
			AssertEqual(e.Commands.GetLowVoltage, true)
			AssertEqual(e.Commands.SetMaxVoltage, float32(3.6))
			AssertEqual(e.Commands.SetMinVoltage, float32(2.496))
			AssertEqual(e.Commands.SetOverVoltage, float32(3.648))
			AssertEqual(e.Commands.GetStatus, true)
			AssertEqual(e.Commands.GetAddrTemp, 1)
			AssertEqual(e.Commands.SetFanMaxTemp, 151)
			AssertEqual(e.Commands.SetStopDissipatingTemp, 171)
			AssertEqual(e.Commands.SetFanLowTemp, 120)
			AssertEqual(e.Commands.GetCellsTemp, true)
		})

		// It("should return an ExecutorCommandValues object got values", func() {
		// 	e := NewExecutor(NewMk3DT(NewMockPort()), "./fixtures/get_commands.yaml", -1)
		// 	e.Run()
		// 	AssertEqual(false, true)
		// })

		// It("should return an ExecutorCommandValues object with set values", func() {
		// 	e := NewExecutor(NewMk3DT(NewMockPort()), "./fixtures/set_commands.yaml", -1)
		// 	e.Run()
		// 	AssertEqual(false, true)
		// })
	})

	Report(t)
}
