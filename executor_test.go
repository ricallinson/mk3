package main

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestExecutor(t *testing.T) {

	Describe("Executor()", func() {

		It("should return a Executor object", func() {
			e := NewExecutor(NewMk3DT(NewMockPort()), "./fixtures/all_commands.yaml")
			AssertEqual(reflect.TypeOf(e).String(), "*main.Executor")
		})
	})

	Report(t)
}
