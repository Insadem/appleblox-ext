//go:build darwin
// +build darwin

package backup

import (
	"os"
	"testing"
)

func TestBackup(t *testing.T) {
	backup, err := New()
	if err != nil {
		t.Error(err)
		return
	}

	appPath := backup.AppPath()
	if _, err := os.Stat(appPath); err != nil {
		t.Error(err)
		return
	}

	err = backup.Close()
	if err != nil {
		t.Error(err)
		return
	}

	if _, err := os.Stat(appPath); err == nil {
		t.Error("Roblox app still exists")
		return
	}
}
