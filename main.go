package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func usage() {
	fmt.Printf(`%s </path/to/keyboard/device> <mouse_module_name>

Disable mouse for 2 seconds while typing on keyboard. This program must be ran as root.
More information: https://github.com/oelmekki/grumpy_cat
`, os.Args[0])
}

var device string
var module string
var lastActivityTimeStamp int64
var timeoutNanoseconds int64 = 750000
var mouseActive = true
var signals chan os.Signal
var exiting bool

func main() {
	if len(os.Args) != 3 {
		usage()
		os.Exit(1)
	}

	device = os.Args[1]
	module = os.Args[2]

	ping := make(chan bool)
	signals = make(chan os.Signal)

	go startPollingKeyboard(ping)

	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go cleanup()

	for {
		select {
		case <-ping:
			if !exiting {
				lastActivityTimeStamp = time.Now().UnixNano()
				mouseActive = false
				deactivateMouse()
			}

		case <-time.After(time.Millisecond * 250):
			if !mouseActive && lastActivityTimeStamp+timeoutNanoseconds < time.Now().UnixNano() {
				mouseActive = true
				activateMouse()
			}
		}
	}
}

func cleanup() {
	<-signals
	exiting = true
	if !mouseActive {
		activateMouse()
		os.Exit(130)
	}
}
