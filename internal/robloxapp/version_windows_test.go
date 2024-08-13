package robloxapp

import "testing"

func TestClientVersion(t *testing.T) {
	version, err := clientVersion()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("client version:", version)
}
