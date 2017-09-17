package main

import (
	"fmt"
	"os/exec"
)

func deactivateMouse() {
	cmd := exec.Command("modprobe", "-r", module)
	_, err := cmd.CombinedOutput()
	if err != nil && !exiting {
		fmt.Printf("modprobe -r %s\n", module)
		fmt.Printf("Can't unload module: %s\n", err)
	}
}
