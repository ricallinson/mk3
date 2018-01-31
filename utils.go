package main

func copyIntoArray(s []byte, d []byte) int {
	for i, _ := range d {
		if i >= len(s) || i >= len(d) {
			return i
		}
		d[i] = s[i]
	}
	return len(d)
}

// Convert a boolean to an integer of 0 or 1.
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
