package syncbreaker

import (
	"insadem/multi-roblox/internal/robloxapp"
	"testing"
	"time"
)

func TestBreak(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Error(err)
	}

	close, err := robloxapp.Open()
	if err != nil {
		t.Error(err)
	}
	defer close()
	time.Sleep(time.Millisecond * 666)

	err = b.Break()
	if err != nil {
		t.Error(err)
	}

	err = b.Break()
	if err == nil {
		t.Error("expected to be not able destroy semaphore")
	}
}
