package main

import (
	"fmt"
	"os/exec"
)

func activateMouse() {
	cmd := exec.Command("modprobe", module)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("modprobe %s\n", module)
		fmt.Printf("Can't load module: %s\n", err)
	}
}
