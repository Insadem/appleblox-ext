package robloxapp

import (
	"encoding/json"
	"net/http"
)

type version struct {
	Value string `json:"clientVersionUpload"`
}

// SCREW YOU WINDOWS, I HAD TO IMPLEMENT VERSION PARSING JUST TO OPEN ROBLOX APP!!!??
func clientVersion() (string, error) {
	url := "https://clientsettingscdn.roblox.com/v2/client-version/WindowsPlayer/channel/LIVE"

	r, err := http.Get(url)
	if err != nil {
		return "", err
	}

	var v version
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return "", err
	}

	return v.Value, nil
}
