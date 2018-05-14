package main

import (
	"fmt"
	"os"
	"strconv"
)

func readInput() string {
	var r string
	fmt.Scanln(&r)
	return r
}

func findFirstAddr(mk3DT *Mk3DT) (int, int, int) {
	for addr := 0; addr <= 255; addr++ {
		if size := mk3DT.GetNumCells(addr); size > 0 {
			return addr, mk3DT.GetSerialNum(addr), size
		}
		fmt.Print(".")
	}
	fmt.Println("")
	return 0, 0, 0
}

// Walk the bus from 1 to 255 and validate the found serial number against the expected.
// Map of Serial Numbers and cell count reaming.
func checkBus(mk3DT *Mk3DT, cards map[int]int) bool {
	// Sum of the maps values.
	totalAddrs := 0
	for v, _ := range cards {
		totalAddrs = totalAddrs + v
	}
	for addr := 0; addr <= 255; addr++ {
		sn := mk3DT.GetSerialNum(addr)
		if sn == 0 {
			// Nothing found at address.
			fmt.Print(".")
			continue
		}
		if _, ok := cards[sn]; !ok {
			// The serial number found was not in the in the given map.
			fmt.Printf("\n\rSerial number '%50d' was found on the BMS bus but not registered in the setup process.\n\r", sn)
			return false
		}
		if cards[sn] <= 0 {
			// The serial number has already had it's cells accounted for.
			fmt.Printf("\n\rSerial number '%05d' has miss matched cells.\n\r", sn)
			return false
		}
		if cards[sn] > 0 {
			//
			fmt.Printf("\n\r%03d S/N: %05d", addr, sn)
			cards[sn]-- // Reduce the cell count by one.
		}
	}
	return true
}

// Prints to standard out the result.
func setupBus(mk3DT *Mk3DT) int {
	fmt.Println("Each BMS card must be setup individually before they can be connected together.")
	fmt.Println("Enter the number of BMS cards you are using.")
	cards, _ := strconv.Atoi(readInput())
	if cards < 1 || cards > 255 {
		fmt.Println("There must be at least one card but no more than 255 in the BMS.")
		return 1
	}
	// Scan the bus for cards.
	used := map[int]int{}
	current := 1
	for cards > 0 {
		fmt.Println("Remove all BMS cards and attach the next BMS card to the Dongle Terminator then press enter.")
		readInput()
		fmt.Println("Scanning Dongle Terminator for BMS card.")
		addr, sn, size := findFirstAddr(mk3DT)
		if used[sn] > 0 {
			fmt.Println("This BMS card has already been used.")
			continue
		}
		if addr > 0 && mk3DT.ChangeAddr(addr, current) {
			fmt.Printf("BMS card address was changed from %03d to %03d.\n\r", addr, current)
			current = current + size
			used[sn] = size
			cards--
		} else {
			if addr == 0 {
				fmt.Println("Could not find a BMS card. Please start again.")
				os.Exit(1)
			}
			fmt.Println("Error setting the new address for the attached BMS card. Please start again.")
			os.Exit(1)
		}
	}
	fmt.Println("Attach all the BMS cards to the Dongle Terminator then press enter.")
	readInput()
	if !checkBus(mk3DT, used) {
		fmt.Println("Error scanning bus. Please start again.")
		return 1
	}
	fmt.Println("\n\rSetup complete.")
	return 0
}
