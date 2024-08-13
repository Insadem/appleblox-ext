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

func libDirPath(path string) (string, error) {
	ld := findLibDir(path)
	if ld == "" {
		return "", os.ErrNotExist
	}

	dir := "/" + ld
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		return "", err
	}

	return dir, nil
}

var LibDir = New(func() (string, error) {
	pth, err := os.Executable()
	if err != nil {
		return "", err
	}

	pth, err = libDirPath(pth)
	if err != nil {
		// Try fallback to current work directory, for example when testing
		pth, err = os.Getwd()
		if err != nil {
			return "", err
		}

		pth, err = libDirPath(pth)
		if err != nil {
			return "", err
		}

		return pth, nil
	}

	return pth, nil
})
