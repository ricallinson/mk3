package main

import (
	"fmt"
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
