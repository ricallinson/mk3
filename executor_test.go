package main

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestExecutor(t *testing.T) {

	var e *Executor

	BeforeEach(func() {
		e = NewExecutor(NewMk3DT(NewMockPort()))
	})

	AfterEach(func() {
		e.Close()
	})

	Describe("Executor()", func() {

		It("should return an Executor object", func() {
			AssertEqual(reflect.TypeOf(e).String(), "*main.Executor")
		})

		It("should populate the ExecutorCommands object from YAML", func() {
			e := NewExecutor(NewMk3DT(NewMockPort()))
			e.Commands = readYamlFileToExecutorCommands("./fixtures/all_commands.yaml")
			AssertEqual(e.Commands.SetStopChargeTemp, 120)
			AssertEqual(e.Commands.GetStopChargeTemp, true)
			AssertEqual(e.Commands.DisableStopChargeTemp, true)
			AssertEqual(e.Commands.DisableShunt, true)
			AssertEqual(e.Commands.EnableShunt, true)
			AssertEqual(e.Commands.SetForceFan, 4)
			AssertEqual(e.Commands.GetHighVoltage, true)
			AssertEqual(e.Commands.ClearMaxVoltageHistory, true)
			AssertEqual(e.Commands.ClearMinVoltageHistory, true)
			AssertEqual(e.Commands.ClearVoltageHistory, true)
			AssertEqual(e.Commands.TriggerLights, true)
			AssertEqual(e.Commands.GetMaxVoltageDetected, true)
			AssertEqual(e.Commands.GetMinVoltageDetected, true)
			AssertEqual(e.Commands.GetStopChargeUnderVoltage, true)
			AssertEqual(e.Commands.SetStopChargeUnderVoltage, 0)
			AssertEqual(e.Commands.GetRealTimeVoltage, true)
			AssertEqual(e.Commands.GetLowVoltage, true)
			AssertEqual(e.Commands.SetMaxVoltage, float32(3.6))
			AssertEqual(e.Commands.SetMinVoltage, float32(2.496))
			AssertEqual(e.Commands.SetOverVoltage, float32(3.648))
			AssertEqual(e.Commands.GetAddrTemp, true)
			AssertEqual(e.Commands.SetFanMaxTemp, 151)
			AssertEqual(e.Commands.SetStopDissipatingTemp, 171)
			AssertEqual(e.Commands.SetFanLowTemp, 120)
			AssertEqual(e.Commands.GetCellsTemp, true)
		})

		It("should return command SetStopChargeTemp", func() {
			e.Commands.SetStopChargeTemp = 116
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetStopChargeTemp", r.Commands), true)
		})

		It("should return '180' from GetStopChargeTemp", func() {
			e.Commands.GetStopChargeTemp = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(r.StopChargeTemp, 180)
		})

		It("should return command DisableStopChargeTemp", func() {
			e.Commands.DisableStopChargeTemp = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("DisableStopChargeTemp", r.Commands), true)
		})

		It("should return command DisableShunt", func() {
			e.Commands.DisableShunt = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("DisableShunt", r.Commands), true)
		})

		It("should return command EnableShunt", func() {
			e.Commands.EnableShunt = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("EnableShunt", r.Commands), true)
		})

		It("should return command SetForceFan", func() {
			e.Commands.SetForceFan = 4
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetForceFan", r.Commands), true)
		})

		It("should return '3.9' from GetHighVoltage", func() {
			e.Commands.GetHighVoltage = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(r.HighVoltage, float32(3.9))
		})

		It("should return command ClearMaxVoltageHistory", func() {
			e.Commands.ClearMaxVoltageHistory = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("ClearMaxVoltageHistory", r.Commands), true)
		})

		It("should return command ClearMinVoltageHistory", func() {
			e.Commands.ClearMinVoltageHistory = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("ClearMinVoltageHistory", r.Commands), true)
		})

		It("should return command ClearVoltageHistory", func() {
			e.Commands.ClearVoltageHistory = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("ClearVoltageHistory", r.Commands), true)
		})

		It("should return command TriggerLights", func() {
			e.Commands.TriggerLights = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("TriggerLights", r.Commands), true)
		})

		It("should return '3.971' from GetMaxVoltage", func() {
			e.Commands.GetMaxVoltageDetected = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(r.MaxVoltageDetected, float32(3.971))
		})

		It("should return '2.432' from GetMinVoltage", func() {
			e.Commands.GetMinVoltageDetected = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(r.MinVoltageDetected, float32(2.432))
		})

		It("should return 'true' from GetStopChargeUnderVoltage", func() {
			e.Commands.GetStopChargeUnderVoltage = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(r.StopChargeUnderVoltage, false)
		})

		It("should return command SetStopChargeUnderVoltage -1", func() {
			e.Commands.SetStopChargeUnderVoltage = -1
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetStopChargeUnderVoltage", r.Commands), true)
		})

		It("should return command SetStopChargeUnderVoltage 2", func() {
			e.Commands.SetStopChargeUnderVoltage = 2
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetStopChargeUnderVoltage", r.Commands), false)
		})

		It("should return command SetStopChargeUnderVoltage 0", func() {
			e.Commands.SetStopChargeUnderVoltage = 0
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetStopChargeUnderVoltage", r.Commands), false)
		})

		It("should return command SetStopChargeUnderVoltage 1", func() {
			e.Commands.SetStopChargeUnderVoltage = 1
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetStopChargeUnderVoltage", r.Commands), true)
		})

		It("should return '3.4' from GetRealTimeVoltage", func() {
			e.Commands.GetRealTimeVoltage = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(r.RealTimeVoltage, float32(3.4))
		})

		It("should return '2.432' from GetLowVoltage", func() {
			e.Commands.GetLowVoltage = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(r.LowVoltage, float32(2.432))
		})

		It("should return command SetMaxVoltage", func() {
			e.Commands.SetMaxVoltage = 3.654
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetMaxVoltage", r.Commands), true)
		})

		It("should return command SetMinVoltage", func() {
			e.Commands.SetMinVoltage = 2.9
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetMinVoltage", r.Commands), true)
		})

		It("should return command SetOverVoltage", func() {
			e.Commands.SetOverVoltage = 3.7
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetOverVoltage", r.Commands), true)
		})

		It("should return '120' from GetAddrTemp", func() {
			e.Commands.GetAddrTemp = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(r.AddrTemp, 120)
		})

		It("should return command SetFanMaxTemp", func() {
			e.Commands.SetFanMaxTemp = 151
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetFanMaxTemp", r.Commands), true)
		})

		It("should return command SetStopDissipatingTemp", func() {
			e.Commands.SetStopDissipatingTemp = 171
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetStopDissipatingTemp", r.Commands), true)
		})

		It("should return command SetFanLowTemp", func() {
			e.Commands.SetFanLowTemp = 120
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(stringInSlice("SetFanLowTemp", r.Commands), true)
		})

		It("should return '120' from GetCellsTemp", func() {
			e.Commands.GetCellsTemp = true
			r := e.ExecuteCommandsAtAddr(0)
			AssertEqual(r.CellsTemp, 0)
		})

		It("should return '255' from ExecuteCommands", func() {
			r := e.ExecuteCommands(255)
			AssertEqual(len(r), 256)
		})
	})

	Report(t)
}
