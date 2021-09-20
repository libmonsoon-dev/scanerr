package scanerr

import (
	"golang.org/x/tools/go/packages"

	"github.com/libmonsoon-dev/scanerr/internal/source"
)

type SourceStringsExtractor interface {
	ExtractStrings(pks []*packages.Package) (result []source.String)
}
