# mk3

[![Build Status](https://travis-ci.org/ricallinson/mk3.svg?branch=master)](https://travis-ci.org/ricallinson/mk3) [![Build status](https://ci.appveyor.com/api/projects/status/fukrjc3xponxntry/branch/master?svg=true)](https://ci.appveyor.com/project/ricallinson/mk3/branch/master)

__UNSTABLE__

Command line interface for the Manzanita Micro USB Dongle Terminator. This is an alternative to the [MK3 Digital Perl Scanner Software](http://www.manzanitamicro.com/downloads/category/5-bms2?download=93%3Aperlscanner) provided by Manzanita Micro.

## Usage

Requires a [Go](https://golang.org/dl/) environment.

    go get github.com/ricallinson/mk3
    go install github.com/ricallinson/mk3

## Examples

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

Changes the address of the first card found to the one given. This should be used when only one card is attached to the dongle.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -new-addr=5

### Bus Address to Execute Commands

A single bus address to target. The bus address is the number of the first cell a card is reasonable for for.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./fixtures/trigger_lights.yaml -addr 1

### Send Raw Command

Send a command as detailed in the [MK3 Digital Perl User Manual Command List](http://www.manzanitamicro.com/downloads/category/5-bms2?download=93%3Aperlscanner).

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -raw "XXX"

### Maximum Cell Address

Option to limit the number of cells that will be addressed on the bus. This is useful to stop the program early when you know the maximum number of cells.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./fixtures/get_commands.yaml -max-addr=50 

### Scan Bus for Cards

Scans the bus for all cards connected.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -scan-cards
	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -scan-cells -max-addr=50 

### Scan Bus for Cells

Scans the bus for all cells on connected cards.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -scan-cells
	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -scan-cells -max-addr=50 

### Realtime

Prints a JSON object. This command loops once it reaches the last cell and is useful for monitoring when setting up a pack or as a log output.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -realtime
	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -realtime -max-addr=68

## Setup Development Environment

Requires a [Go](https://golang.org/dl/) environment.

	git clone git@github.com:ricallinson/mk3.git $GOPATH/src/git@github.com/ricallinson/mk3
    go get github.com/tarm/serial
    go get gopkg.in/yaml.v2
    go get github.com/ricallinson/simplebdd
    cd $GOPATH/src/git@github.com/ricallinson/mk3
    go install

## Testing

	cd $GOPATH/src/git@github.com/ricallinson/mmz
	go test

## Coverage

	cd $GOPATH/src/git@github.com/ricallinson/mmz
	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out
