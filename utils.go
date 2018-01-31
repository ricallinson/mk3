package main

import (
	"fmt"
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
