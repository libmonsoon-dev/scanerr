package config

import "runtime"

type ScannerConfig struct {
	PackagesLoaderConf LoaderConfig
}

type AppConfig struct {
	StringsExtractorConfig StringsExtractorConfig
	CacheConfig            CacheConfig
}

func DefaultScannerConfig() ScannerConfig {
	return ScannerConfig{
		PackagesLoaderConf: LoaderConfig{
			Patterns: []string{"./..."},
		},
	}
}

func DefaultAppConfig() AppConfig {
	return AppConfig{
		StringsExtractorConfig: StringsExtractorConfig{
			NumWorkers: runtime.NumCPU(),
		},
		CacheConfig: CacheConfig{
			UseLoaderCache: true,
		},
	}
}
