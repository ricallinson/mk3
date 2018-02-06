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
	for addr := 1; addr <= 255; addr++ {
		if size := mk3DT.GetNumCells(addr); size > 0 {
			return addr, mk3DT.GetSerialNum(addr), size
		}
		fmt.Print(".")
	}
	fmt.Println("")
	return 0, 0, 0
}

func checkBus(mk3DT *Mk3DT, cards map[int]int) bool {
	for addr := 1; addr <= 255; addr++ {
		sn := mk3DT.GetSerialNum(addr)
		if cards[sn] > 0 {
			fmt.Println(addr, sn)
			cards[sn]--
		} else if sn > 0 && cards[sn] <= 0 {
			return false
		} else {
			fmt.Print(".")
		}
	}
	return true
}

// Prints to standard out the result.
func setupBus(mk3DT *Mk3DT) int {
	fmt.Println("Enter the number of MK3 cards you are using for your BMS.")
	cards, _ := strconv.Atoi(readInput())
	if cards < 1 || cards > 255 {
		fmt.Println("There must be at least one card but no more than 255 in the BMS.")
		return 1
	}
	// Scan the bus for cards.
	used := map[int]int{}
	current := 1
	for cards > 0 {
		fmt.Println("Remove the current BMS card and attach the next BMS card to the Dongle Terminator then press enter.")
		readInput()
		fmt.Println("Scanning Dongle Terminator for BMS card.")
		addr, sn, size := findFirstAddr(mk3DT)
		if used[sn] > 0 {
			fmt.Println("This BMS card has already been used.")
			continue
		}
		if addr > 0 && mk3DT.ChangeAddr(addr, current) {
			fmt.Printf("BMS card address was changed from %v to %v.\n\r", addr, current)
			current = current + size
			used[sn] = size
			cards--
		} else {
			if addr == 0 {
				fmt.Println("Could not find BMS card. Please start again.")
				os.Exit(1)
			}
			fmt.Println("Error setting new address for BMS card. Please start again.")
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
