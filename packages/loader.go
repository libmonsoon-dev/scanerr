package packages

import (
	"fmt"

	"github.com/libmonsoon-dev/scanerr/errutils"

	"golang.org/x/tools/go/packages"
)

type Loader interface {
	Load() ([]*packages.Package, error)
}

type LoaderConfig struct {
	UseCache bool
	Patterns []string
}

func NewLoader(conf LoaderConfig) (l Loader) {
	l = &loader{
		cfg: &packages.Config{
			Mode: packages.LoadSyntax, // TODO: LoadAllSyntax
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

	fmt.Printf("loaded %v packages\n", len(pkgs))

	return pkgs, nil
}

var _ Loader = (*loader)(nil)
