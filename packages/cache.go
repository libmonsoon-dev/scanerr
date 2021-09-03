package packages

import (
	"sync"

	"golang.org/x/tools/go/packages"
)

type loaderCache struct {
	loader     Loader
	loaderOnce sync.Once

	result []*packages.Package
	error  error
}

func (lc *loaderCache) Load() ([]*packages.Package, error) {
	lc.loaderOnce.Do(func() {
		lc.result, lc.error = lc.loader.Load()
	})

	return lc.result, lc.error
}

var _ Loader = (*loaderCache)(nil)
