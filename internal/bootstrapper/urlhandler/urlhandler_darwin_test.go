package urlhandler

import "testing"

func TestUrlHandler(t *testing.T) {
	h, err := New()
	if err != nil {
		t.Error(err)
		return
	}

	v := h.Check(ROBLOX_BUNDLE_IDENTIFIER, "roblox")
	if v != true {
		t.Error("roblox handler supposed to be set to default")
	}

	v = h.Check(ROBLOX_BUNDLE_IDENTIFIER, "roblox-player")
	if v != true {
		t.Error("roblox-player handler supposed to be set to default")
	}

	v = h.Check(ROBLOX_BUNDLE_IDENTIFIER, "non-existant")
	if v == true {
		t.Error("supposed to be not set")
	}
}
