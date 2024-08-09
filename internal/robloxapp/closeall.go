package robloxapp

import (
	"insadem/multi-roblox/pkg/ps"
	"os"
)

func CloseAll() {
	processes, err := ps.Processes()
	if err != nil {
		return
	}

	for _, process := range processes {
		pName := process.Executable()
		if pName == "RobloxPlayer" {
			p, err := os.FindProcess(process.Pid())
			if err == nil {
				p.Kill()
			}
		}
	}
}
