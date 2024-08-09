//go:build darwin
// +build darwin

package backup

import "os/exec"

func New() (Backup, error) {
	rp := "/Applications/Roblox.app"
	dd, err := destDir()
	if err != nil {
		return Backup{}, err
	}

	err = exec.Command("cp", "-a", rp, dd).Run()
	if err != nil {
		return Backup{}, err
	}

	return Backup{dir: dd}, nil
}
