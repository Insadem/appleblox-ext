package syncbreaker

import "testing"

func TestBreak(t *testing.T) {
	err := Break()
	if err != nil {
		t.Error(err)
	}
}
