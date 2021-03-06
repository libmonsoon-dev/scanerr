package scanerr_test

import (
	"testing"

	app "github.com/libmonsoon-dev/scanerr"
	"github.com/libmonsoon-dev/scanerr/config"
)

func Benchmark(b *testing.B) {
	appConf := config.DefaultAppConfig()
	s := app.NewScanerr(appConf)

	scannerConf := config.DefaultScannerConfig()
	scannerConf.PackagesLoaderConf.Patterns = []string{"./testdata/file-not-found/cmd"}

	inputErr := "runtime error: open /not-exist: open /not-exist: file does not exist"
	_, err := s.Scan(inputErr, scannerConf)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := s.Scan(inputErr, scannerConf)
		if err != nil {
			b.Fatal(err)
		}
	}
}
