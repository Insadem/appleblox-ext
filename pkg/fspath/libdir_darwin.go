package fspath

import (
	"insadem/multi-roblox/pkg/fspath/shortenpath"
	"os"
)

// findLibDir returns path to lib dir or empty string.
func findLibDir(path string) string {
	for {
		pth, v := shortenpath.Shorten(path)
		if pth == "" {
			return ""
		}
		if v == "appleblox-ext" {
			return pth + "/appleblox-ext/lib"
		}

		path = pth
	}
}

var LibDir = New(func() (string, error) {
	pth, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir := "/" + findLibDir(pth)
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		return "", err
	}

	return dir, nil
})
