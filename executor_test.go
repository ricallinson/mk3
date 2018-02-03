package main

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestExecutor(t *testing.T) {

	var e *Executor

	BeforeEach(func() {
		e = NewExecutor(NewMk3DT(NewMockPort()), "")
	})

	Describe("Executor()", func() {

		It("should return an Executor object", func() {
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
			AssertEqual(e.Commands.SetFirstPosition, -1)
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

		It("should return 'true' from SetStopChargeTemp", func() {
			e.Commands.SetStopChargeTemp = 116
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.SetStopChargeTemp, true)
		})

		It("should return '180' from GetStopTemp", func() {
			e.Commands.GetStopTemp = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.GetStopTemp, 180)
		})

		It("should return 'true' from DisableStopChargeTemp", func() {
			e.Commands.DisableStopChargeTemp = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.DisableStopChargeTemp, true)
		})

		It("should return 'true' from ChangeAddr", func() {
			e.Commands.ChangeAddr = 3
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.ChangeAddr, true)
		})

		It("should return 'Mk3DTCommands' from GetCommands", func() {
			e.Commands.GetCommands = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(reflect.TypeOf(r.GetCommands).String(), "main.Mk3DTCommands")
		})

		It("should return 'true' from DisableShunt", func() {
			e.Commands.DisableShunt = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.DisableShunt, true)
		})

		It("should return 'true' from EnableShunt", func() {
			e.Commands.EnableShunt = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.EnableShunt, true)
		})

		It("should return 'true' from ForceFan", func() {
			e.Commands.ForceFan = 4
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.ForceFan, true)
		})

		It("should return 'true' from GetFirstPosition", func() {
			e.Commands.GetFirstPosition = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.GetFirstPosition, false)
		})

		It("should return 'true' from SetFirstPosition", func() {
			e.Commands.SetFirstPosition = 1
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.SetFirstPosition, true)
		})

		It("should return '3.9' from GetHighVoltage", func() {
			e.Commands.GetHighVoltage = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.GetHighVoltage, float32(3.9))
		})

		It("should return 'true' from ClearMaxVoltageHistory", func() {
			e.Commands.ClearMaxVoltageHistory = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.ClearMaxVoltageHistory, true)
		})

		It("should return 'true' from ClearMinVoltageHistory", func() {
			e.Commands.ClearMinVoltageHistory = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.ClearMinVoltageHistory, true)
		})

		It("should return 'true' from ClearVoltageHistory", func() {
			e.Commands.ClearVoltageHistory = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.ClearVoltageHistory, true)
		})

		It("should return 'Mk3DTLightsStatus' from TriggerLights", func() {
			e.Commands.TriggerLights = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(reflect.TypeOf(r.TriggerLights).String(), "main.Mk3DTLightsStatus")
		})

		It("should return '3.971' from GetMaxVoltage", func() {
			e.Commands.GetMaxVoltage = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.GetMaxVoltage, float32(3.971))
		})

		It("should return '2.432' from GetMinVoltage", func() {
			e.Commands.GetMinVoltage = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.GetMinVoltage, float32(2.432))
		})

		It("should return 'true' from GetStopChargeUnderVoltage", func() {
			e.Commands.GetStopChargeUnderVoltage = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.GetStopChargeUnderVoltage, false)
		})

		It("should return '3.4' from GetRealTimeVoltage", func() {
			e.Commands.GetRealTimeVoltage = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.GetRealTimeVoltage, float32(3.4))
		})

		It("should return '2.432' from GetLowVoltage", func() {
			e.Commands.GetLowVoltage = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.GetLowVoltage, float32(2.432))
		})

		It("should return 'true' from SetMaxVoltage", func() {
			e.Commands.SetMaxVoltage = 3.654
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.SetMaxVoltage, true)
		})

		It("should return 'true' from SetMinVoltage", func() {
			e.Commands.SetMinVoltage = 2.9
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.SetMinVoltage, true)
		})

		It("should return 'true' from SetOverVoltage", func() {
			e.Commands.SetOverVoltage = 3.7
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.SetOverVoltage, true)
		})

		It("should return 'Mk3DTStatus' from GetStatus", func() {
			e.Commands.GetStatus = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(reflect.TypeOf(r.GetStatus).String(), "main.Mk3DTStatus")
		})

		It("should return '120' from GetAddrTemp", func() {
			e.Commands.GetAddrTemp = 1
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.GetAddrTemp, 120)
		})

		It("should return 'true' from SetFanMaxTemp", func() {
			e.Commands.SetFanMaxTemp = 151
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.SetFanMaxTemp, true)
		})

		It("should return 'true' from SetStopDissipatingTemp", func() {
			e.Commands.SetStopDissipatingTemp = 171
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.SetStopDissipatingTemp, true)
		})

		It("should return 'true' from SetFanLowTemp", func() {
			e.Commands.SetFanLowTemp = 120
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.SetFanLowTemp, true)
		})

		It("should return '120' from GetCellsTemp", func() {
			e.Commands.GetCellsTemp = true
			r := e.ExecuteCommandsOnAddr(0)
			AssertEqual(r.GetCellsTemp, "Cold")
		})
	})

	Report(t)
}
