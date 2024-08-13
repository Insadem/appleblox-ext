package syncbreaker

import (
	"errors"
	"insadem/multi-roblox/pkg/fspath"

	"github.com/ebitengine/purego"
)

type calldef func() int
type Breaker struct {
	destroySemaphore calldef
}

func New() (Breaker, error) {
	dir, err := fspath.LibDir.Get()
	if err != nil {
		return Breaker{}, err
	}

	lib, err := purego.Dlopen(dir+"/syncbreaker_darwin.dylib", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return Breaker{}, err
	}

	var destroySemaphore calldef
	purego.RegisterLibFunc(&destroySemaphore, lib, "destroySemaphore")

	return Breaker{
		destroySemaphore: destroySemaphore,
	}, nil
}

func (b Breaker) Break() error {
	if b.destroySemaphore() != 0 {
		return errors.New("couldn't destroy semaphore")
	}

	return nil
}
