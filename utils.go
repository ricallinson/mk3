package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func copyIntoArray(s []byte, d []byte) int {
	for i, _ := range d {
		if i >= len(s) || i >= len(d) {
			return i
		}
		d[i] = s[i]
	}
	return len(d)
}

// Pad an int to three places and returns as string.
func padInt(value int) string {
	return fmt.Sprintf("%03d", value)
}

// Converts "000F" to an int.
func tempToInt(s string) int {
	t, _ := strconv.ParseInt(strings.TrimSuffix(s, "F"), 10, 32)
	return int(t)
}

// Converts "0.000V" to a float32.
func voltToFloat32(s string) float32 {
	f, _ := strconv.ParseFloat(strings.TrimSuffix(s, "V"), 32)
	return float32(f)
}

// Reads a file into a byte array or exits.
func readFileToByteArray(p string) []byte {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalf("Error reading file: #%v ", err)
	}
	return b
}

func readYamlFileToExecutorCommands(p string) *ExecutorCommands {
	r := &ExecutorCommands{}
	if p != "" {
		err := yaml.Unmarshal(readFileToByteArray(p), r)
		if err != nil {
			log.Fatalf("YAML Parse Error: %v", err)
		}
	}
	return r
}

func interfaceToYaml(i interface{}) []byte {
	d, err := yaml.Marshal(i)
	if err != nil {
		log.Fatalf("YAML Write Error: %v", err)
	}
	return d
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
