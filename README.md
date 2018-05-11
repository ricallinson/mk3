# mk3

Command line interface for the Manzanita Micro USB Dongle Terminator. This is an alternative to the [MK3 Digital Perl Scanner Software](http://www.manzanitamicro.com/downloads/category/5-bms2?download=93%3Aperlscanner) provided by Manzanita Micro.

## Setup Development Environment

Requires a [Go](https://golang.org/dl/) environment.

    go get github.com/tarm/serial
    go get gopkg.in/yaml.v2
    go get github.com/ricallinson/simplebdd

## Examples

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -raw "01l."

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./fixtures/get_commands.yaml

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./fixtures/trigger_lights.yaml -addr 1

## Options

### Dongle Location

__Required__

The path to the USB port.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7

### Path to Commands File

The path to the file containing the commands to execute against the bus.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./fixtures/get_commands.yaml

### Bus Address to Execute Commands

A single bus address to target. The bus address is the number of the first cell a card is reasonable for for.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./fixtures/trigger_lights.yaml -addr 1

### Send Raw Command

Send a command as detailed in the [MK3 Digital Perl User Manual](http://www.manzanitamicro.com/downloads/category/5-bms2?download=93%3Aperlscanner).

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -raw "01l."

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

### Change Cards Bus Address

Changes the address of the first card found to the one given. This should be used when only one card is attached to the dongle.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -new-addr=5

### List Cell Volts

Lists the current voltage of all cells found.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -volts
	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -volts -max-addr=50 

### List Cell Temperatures

Lists the current temperature of all cells found.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -temps
	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -temps -max-addr=50

### Setup

The setup option walks through assigning addresses to each card that will be used in a bus and then validates the setup was successful.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -setup


## Testing

	go test

## Coverage

	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out

