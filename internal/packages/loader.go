package packages

import (
	"fmt"

	"golang.org/x/tools/go/packages"

	"github.com/libmonsoon-dev/scanerr/internal/errutils"
)

type Loader interface {
	Load() ([]*packages.Package, error)
}

type LoaderConfig struct {
	UseCache bool
	Patterns []string
}

func NewLoader(conf LoaderConfig) (l Loader) {
	const mode = packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax

	l = &loader{
		cfg: &packages.Config{
			Mode: mode,
		},
		patterns: conf.Patterns,
	}

	if conf.UseCache {
		l = &loaderCache{loader: l}
	}

	return l
}

type loader struct {
	cfg      *packages.Config
	patterns []string
}

func (l loader) Load() ([]*packages.Package, error) {
	patterns := make([]string, len(l.patterns))
	for i := range l.patterns {
		patterns[i] = fmt.Sprintf("pattern=%v", l.patterns[i])
	}

	pkgs, err := packages.Load(l.cfg, patterns...)
	if err != nil {
		return nil, fmt.Errorf("packages.Load: %w", err)
	}

	pkgErrors := errutils.NewBundle()
	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		for _, err := range pkg.Errors {
			pkgErrors.Add(err)
		}
	})

	if !pkgErrors.IsEmpty() {
		return nil, pkgErrors
	}

	return pkgs, nil
}

var _ Loader = (*loader)(nil)
