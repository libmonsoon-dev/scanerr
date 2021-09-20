package config

import (
	"runtime"

	"github.com/libmonsoon-dev/scanerr/internal/packages"
	"github.com/libmonsoon-dev/scanerr/internal/source"
)

type Config struct {
	PackagesLoaderConf     packages.LoaderConfig
	StringsExtractorConfig source.StringsExtractorConfig
}

func DefaultConfig() Config {
	return Config{
		PackagesLoaderConf: packages.LoaderConfig{
			UseCache: true,
			Patterns: []string{"./..."},
		},
		StringsExtractorConfig: source.StringsExtractorConfig{
			NumWorkers: runtime.NumCPU(),
		},
	}
}
