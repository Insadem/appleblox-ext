package fspath

import "testing"

func TestFindLibDir(t *testing.T) {
	dir := findLibDir("usr/appleblox-ext/lib/keep/search/for/the/truth")
	if dir != "usr/appleblox-ext/lib" {
		t.Error("didn't found lib dir, instead found: " + dir)
	}

	dir = findLibDir("there/is/no/truth")
	if dir != "" {
		t.Error("expected to be empty")
	}
}

func TestLibDir(t *testing.T) {
	_, err := LibDir.Get()
	if err != nil {
		t.Error(err)
	}
}
