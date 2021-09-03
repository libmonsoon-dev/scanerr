//+build wireinject

package app

import (
	"github.com/google/wire"

	"github.com/libmonsoon-dev/scanerr/config"
	"github.com/libmonsoon-dev/scanerr/packages"
	"github.com/libmonsoon-dev/scanerr/scanerr"
	"github.com/libmonsoon-dev/scanerr/searcher"
	"github.com/libmonsoon-dev/scanerr/source"
)

func NewScanerr(_ config.Config) *scanerr.Scanerr {
	panic(wire.Build(
		packages.NewLoader,
		source.NewStringsExtractor,
		scanerr.NewScanerr,
		searcher.NewSearcher,
		wire.FieldsOf(new(config.Config), "ScanerrConfig", "PackagesLoaderConf", "StringsExtractorConfig"),
	))
}
