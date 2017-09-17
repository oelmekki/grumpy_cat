/*
 * input_event kernel mapping and handling in go from this file comes from
 * https://github.com/Banrai/PiScan/blob/master/scanner/scanner.go
 * Thus, same licence applies.
 */
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type InputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

const (
	EVENT_BUFFER   = 64
	EVENT_CAPTURES = 16
)

var EVENT_SIZE = int(unsafe.Sizeof(InputEvent{}))

func cantReadKeyboard(err error) {
	fmt.Printf("Can't read keyboard: %s\n", err)
	os.Exit(1)
}

func startPollingKeyboard(ping chan bool) {
	kbd, err := os.Open(device)
	if err != nil {
		cantReadKeyboard(err)
	}

	defer kbd.Close()

	for {
		events, err := readKeyboard(kbd)
		if err != nil {
			cantReadKeyboard(err)
		}

		for i := range events {
			if events[i].Type == 1 && events[i].Value == 1 {
				ping <- true
			}
		}
	}
}

func readKeyboard(dev *os.File) ([]InputEvent, error) {
	events := make([]InputEvent, EVENT_CAPTURES)
	buffer := make([]byte, EVENT_SIZE*EVENT_CAPTURES)
	_, err := dev.Read(buffer)
	if err != nil {
		return events, err
	}
	b := bytes.NewBuffer(buffer)
	err = binary.Read(b, binary.LittleEndian, &events)
	if err != nil {
		return events, err
	}
	// remove trailing structures
	for i := range events {
		if events[i].Time.Sec == 0 {
			events = append(events[:i])
			break
		}
	}
	return events, err
}
