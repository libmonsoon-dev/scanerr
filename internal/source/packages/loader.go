package packages

import (
	"fmt"

	"golang.org/x/tools/go/packages"

	"github.com/libmonsoon-dev/scanerr/config"
	"github.com/libmonsoon-dev/scanerr/internal/errutils"
)

func NewLoader() (l *Loader) {
	l = &Loader{}
	return l
}

type Loader struct {
}

func (l Loader) Load(conf config.LoaderConfig) ([]*packages.Package, error) {
	const mode = packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax

	patterns := make([]string, len(conf.Patterns))
	for i := range conf.Patterns {
		patterns[i] = fmt.Sprintf("pattern=%v", conf.Patterns[i])
	}

	cfg := &packages.Config{
		Mode: mode,
	}

	pkgs, err := packages.Load(cfg, patterns...)
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
