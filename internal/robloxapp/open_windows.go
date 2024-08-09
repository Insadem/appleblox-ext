package robloxapp

import (
	"insadem/multi-roblox/pkg/fspath"
	"os/exec"
)

func exePath() (string, error) {
	hmDir, err := fspath.HomeDir.Get()
	if err != nil {
		return "", err
	}

	v, err := clientVersion()
	if err != nil {
		return "", err
	}

	return hmDir + "\\AppData\\Local\\Roblox\\Versions\\" + v + "\\RobloxPlayerBeta.exe", nil
}

func Open() (func(), error) {
	rbxPath, err := exePath()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(rbxPath)
	cmd.Start()

	return func() {
		cmd.Process.Kill()
		cmd.Process.Release()
	}, nil
}
