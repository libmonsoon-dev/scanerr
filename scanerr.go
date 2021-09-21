//go:build wireinject

package scanerr

import (
	"github.com/google/wire"

	"github.com/libmonsoon-dev/scanerr/internal/source/ast"

	"github.com/libmonsoon-dev/scanerr/config"
	"github.com/libmonsoon-dev/scanerr/internal/cache"
	"github.com/libmonsoon-dev/scanerr/internal/scanerr"
	"github.com/libmonsoon-dev/scanerr/internal/source/packages"
	"github.com/libmonsoon-dev/scanerr/internal/unfmt"
)

func NewScanerr(_ config.AppConfig) *scanerr.Scanner {
	panic(wire.Build(
		scanerr.NewScanner,
		newPackageLoader,
		newPackagesStringsExtractor,
		newASTStringsExtractor,
		unfmt.NewStringsMatcher,

		wire.Bind(new(scanerr.SourceStringMatcher), new(*unfmt.StringsMatcher)),

		wire.FieldsOf(new(config.AppConfig), "StringsExtractorConfig", "CacheConfig"),
	))
}

func newPackageLoader(cacheConfig config.CacheConfig) (l scanerr.PackagesLoader) {
	l = packages.NewLoader()

	if cacheConfig.UseLoaderCache {
		l = cache.NewPackageLoader(l)
	}

	return
}

func newPackagesStringsExtractor(conf config.StringsExtractorConfig, factory packages.ASTStringExtractorFactory) scanerr.SourceStringsExtractor {
	if conf.NumWorkers > 1 {
		return packages.NewConcurrentStringsExtractor(conf, factory)
	}
	return packages.NewStringsExtractor(factory)
}

func newASTStringsExtractor(cacheConfig config.CacheConfig) (f packages.ASTStringExtractorFactory) {
	f = ast.NewStringExtractorFactory()

	if cacheConfig.UseStringsExtractorCache {
		f = cache.NewStringExtractorFactory(f)
	}
	return
}
