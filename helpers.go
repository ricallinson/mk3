package main

import (
	"fmt"
	"log"
)

// Prints to standard out the result.
func sendRawCommand(mk3DT *Mk3DT, s string) {
	log.Println(string(mk3DT.Raw(s + "\n\r")))
}

// Prints to standard out the result.
func scanForCells(mk3DT *Mk3DT, minAddr int, maxAddr int) int {
	pos := 1
	lastSn := 0
	lastEmpty := false
	for addr := 1; addr <= maxAddr; addr++ {
		if sn := mk3DT.GetSerialNum(addr); sn > 0 {
			if sn == lastSn {
				pos++
			} else {
				pos = 1
			}
			if lastEmpty {
				fmt.Println("")
			}
			fmt.Printf("Cell %03d on card %05d at position %03d\n\r", addr, sn, pos)
			lastEmpty = false
			lastSn = sn
		} else {
			lastEmpty = true
			fmt.Print(".")
		}
	}
	fmt.Println("Done")
	return 0
}

func scanForCards(mk3DT *Mk3DT, maxAddr int) int {
	lastSn := 0
	count := 1
	for addr := 1; addr <= maxAddr; addr++ {
		sn := mk3DT.GetSerialNum(addr)
		if sn != lastSn && sn != 0 {
			cells := mk3DT.GetNumCells(addr)
			fmt.Printf("%02d: S/N %05d with cells %03d-%03d/%d\n\r", count, sn, addr, addr+cells-1, cells)
			lastSn = sn
			count++
		}
	}
	return 0
}

func setAddr(mk3DT *Mk3DT, newAddr int) int {
	for addr := 1; addr <= 255; addr++ {
		fmt.Print(".")
		if sn := mk3DT.GetSerialNum(addr); sn > 0 {
			mk3DT.ChangeAddr(addr, newAddr)
			fmt.Printf("\n\rAddress set to %03d\n\r", newAddr)
			return 0
		}
	}
	fmt.Println("")
	return 1
}
