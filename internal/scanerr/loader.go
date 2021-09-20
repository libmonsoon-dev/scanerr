package scanerr

import (
	"golang.org/x/tools/go/packages"

	"github.com/libmonsoon-dev/scanerr/config"
)

type PackagesLoader interface {
	Load(conf config.LoaderConfig) ([]*packages.Package, error)
}
