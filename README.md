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

## Notes

	mk3 -dongle /dev/tty.usbserial-A904RBQ7
