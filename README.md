# mk3

[![Build Status](https://travis-ci.org/ricallinson/mk3.svg?branch=master)](https://travis-ci.org/ricallinson/mk3) [![Build status](https://ci.appveyor.com/api/projects/status/fukrjc3xponxntry/branch/master?svg=true)](https://ci.appveyor.com/project/ricallinson/mk3/branch/master)

Command line interface for the Manzanita Micro USB Dongle Terminator. This is an alternative to the [MK3 Digital Perl Scanner Software](http://www.manzanitamicro.com/downloads/category/5-bms2?download=93%3Aperlscanner) provided by Manzanita Micro.

## Usage

Download the executable for your chosen platform from the [releases](https://github.com/ricallinson/mk3/releases/tag/v1.0) page. Feel free to rename the executable to `mk3` or `mk3.exe` depending on your chosen platform.

You will need to know the location of the USB serial port which the dongle is plugged into. The [MK3 Digital Perl Scanner Software](http://www.manzanitamicro.com/downloads/category/5-bms2?download=93%3Aperlscanner) documentation describes how to find this for Windows as a COM port number. For Unix based systems you can use [dmesg | grep tty](https://www.cyberciti.biz/faq/find-out-linux-serial-ports-with-setserial/) as described in the link.

### Examples

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -realtime

    mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./samples/get_settings.yaml

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./samples/get_settings.yaml -addr 4

## Options

### Dongle Location (required)

The path to the USB port where the dongle is connected.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7

### Path to Commands File

The path to the file containing the commands to execute against the bus.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./samples/get_settings.yaml

### Change a Cards Bus Address

Changes the address of the first card found to the one given. This should be used when only one card is attached to the dongle. If more than one card is connect the first card found will be given the new address.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -new-addr=5

### Bus Address to Execute Commands

A single bus address to target sending the command file to.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./fixtures/trigger_lights.yaml -addr 1

### Send Raw Command

Send a command as detailed in the [MK3 Digital Perl User Manual Command List](http://www.manzanitamicro.com/downloads/category/5-bms2?download=93%3Aperlscanner).

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -raw "XXX"

### Maximum Cell Address

Option to limit the number of cells that will be addressed on the bus. This is useful to stop the program early when you know the maximum number of cells.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./fixtures/get_commands.yaml -max-addr=68 

### Scan Bus for Cards

Scans the bus for all cards connected.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -scan-cards
	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -scan-cells -max-addr=68 

### Scan Bus for Cells

Scans the bus for all cells on connected cards.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -scan-cells
	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -scan-cells -max-addr=68 

### Realtime

This command loops once it reaches the last cell and is useful for monitoring when setting up a pack or as a log output.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -realtime
	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -realtime -max-addr=68

Outputs YAML to `stdout` with the following structure;

	Timestamp int64
	Address   int
	SerialNum int
	CellCount int
	Volts     float32
	MaxVolts  float32
	MinVolts  float32
	Temp      int

## Setup Development Environment

Requires a [Go](https://golang.org/dl/) environment.

	git clone git@github.com:ricallinson/mk3.git $GOPATH/src/git@github.com/ricallinson/mk3
    cd $GOPATH/src/git@github.com/ricallinson/mk3
    go get ./...
    go install

## Testing

	cd $GOPATH/src/git@github.com/ricallinson/mmz
	go test

## Coverage

	cd $GOPATH/src/git@github.com/ricallinson/mmz
	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out
