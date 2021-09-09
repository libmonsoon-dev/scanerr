//go:build wireinject

package app

import (
	"github.com/google/wire"

	"github.com/libmonsoon-dev/scanerr/config"
	"github.com/libmonsoon-dev/scanerr/packages"
	"github.com/libmonsoon-dev/scanerr/scanerr"
	"github.com/libmonsoon-dev/scanerr/source"
	"github.com/libmonsoon-dev/scanerr/unfmt"
)

func NewScanerr(_ config.Config) *scanerr.Scanerr {
	panic(wire.Build(
		packages.NewLoader,
		source.NewStringsExtractor,
		scanerr.NewScanerr,
		unfmt.NewMatcher,
		wire.FieldsOf(new(config.Config), "PackagesLoaderConf", "StringsExtractorConfig"),
	))
}
