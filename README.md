# mk3
Command line interface for the Manzanita Micro USB Dongle Terminator.

## Setup Development Environment

Requires a [Go](https://golang.org/dl/) environment.

    go get github.com/tarm/serial
    go get gopkg.in/yaml.v2
    go get github.com/ricallinson/simplebdd

## Testing

	go test

## Coverage

	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out

## Examples

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -raw "01l."

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./fixtures/get_commands.yaml

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -cmd ./fixtures/trigger_lights.yaml -addr 1

## Options

### Dongle Location (required)

### Path to Commands File

### Bus Address to Execute Commands

### Send Raw Command

### Scan Bus for Cards

### Maximum Bus Address

### Change Cards Bus Address

### List Cell Volts

### List Cell Temperatures

### Setup

The setup option walks through assigning addressess to eack MK3 card and then validates the setup was successful.

	mk3 -dongle /dev/tty.usbserial-A904RBQ7 -setup

