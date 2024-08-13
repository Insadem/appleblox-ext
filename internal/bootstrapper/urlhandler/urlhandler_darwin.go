package urlhandler

import (
	"sync"

	"github.com/ebitengine/purego"
)

type calldef func(bundleIdentifier, urlScheme string) int

type UrlHandler struct {
	set   calldef
	check calldef
	mut   *sync.Mutex
}

func New() (UrlHandler, error) {
	lib, err := purego.Dlopen("liburlhandler.a", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return UrlHandler{}, err
	}

	var set, check calldef
	purego.RegisterLibFunc(&set, lib, "set")
	purego.RegisterLibFunc(&check, lib, "check")

	return UrlHandler{
		set:   set,
		check: check,
	}, nil
}

func (u UrlHandler) Set(bundleIdentifier, urlScheme string) bool {
	u.mut.Lock()
	defer u.mut.Unlock()

	u.set()
}

func (u UrlHandler) Check(bundleIdentifier, urlScheme string) bool {
	u.mut.Lock()
	defer u.mut.Unlock()

}
