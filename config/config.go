package config

import (
	"runtime"

	"github.com/libmonsoon-dev/scanerr/packages"
	"github.com/libmonsoon-dev/scanerr/scanerr"
	"github.com/libmonsoon-dev/scanerr/source"
)

type Config struct {
	ScanerrConfig          scanerr.Config
	PackagesLoaderConf     packages.LoaderConfig
	StringsExtractorConfig source.StringsExtractorConfig
}

func DefaultConfig() Config {
	return Config{
		ScanerrConfig: scanerr.Config{},
		PackagesLoaderConf: packages.LoaderConfig{
			UseCache: true,
			Patterns: []string{"./..."},
		},
		StringsExtractorConfig: source.StringsExtractorConfig{
			NumWorkers: runtime.NumCPU(),
		},
	}
}
