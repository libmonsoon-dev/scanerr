//go:build wireinject

package scanerr

import (
	"github.com/google/wire"

	"github.com/libmonsoon-dev/scanerr/config"
	"github.com/libmonsoon-dev/scanerr/internal/packages"
	"github.com/libmonsoon-dev/scanerr/internal/scanerr"
	"github.com/libmonsoon-dev/scanerr/internal/source"
	"github.com/libmonsoon-dev/scanerr/internal/unfmt"
)

func NewScanerr(_ config.Config) *scanerr.Scanner {
	panic(wire.Build(
		packages.NewLoader,
		source.NewStringsExtractor,
		scanerr.NewScanner,
		unfmt.NewMatcher,
		wire.FieldsOf(new(config.Config), "PackagesLoaderConf", "StringsExtractorConfig"),
	))
}
