package urlhandler

import (
	"insadem/multi-roblox/pkg/fspath"
	"sync"

	"github.com/ebitengine/purego"
)

const (
	ROBLOX_BUNDLE_IDENTIFIER = "com.roblox.RobloxPlayer"
	APP_BUNDLE_IDENTIFIER    = "" // TODO: define later, maybe take it from .ENV file
)

type calldef func(bundleIdentifier, urlScheme string) int

type UrlHandler struct {
	set   calldef
	check calldef
	mut   *sync.Mutex
}

func New() (UrlHandler, error) {
	dir, err := fspath.LibDir.Get()
	if err != nil {
		return UrlHandler{}, err
	}

	lib, err := purego.Dlopen(dir+"/urlhandler_darwin.dylib", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return UrlHandler{}, err
	}

	var set, check calldef
	purego.RegisterLibFunc(&set, lib, "set")
	purego.RegisterLibFunc(&check, lib, "check")

	return UrlHandler{
		set:   set,
		check: check,
		mut:   &sync.Mutex{},
	}, nil
}

func (u UrlHandler) Set(bundleIdentifier, urlScheme string) bool {
	u.mut.Lock()
	defer u.mut.Unlock()

	if u.set(bundleIdentifier, urlScheme) == 0 {
		return true
	} else {
		return false
	}
}

func (u UrlHandler) Check(bundleIdentifier, urlScheme string) bool {
	u.mut.Lock()
	defer u.mut.Unlock()

	if u.check(bundleIdentifier, urlScheme) == 0 {
		return true
	} else {
		return false
	}
}
