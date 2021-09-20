package scanerr_test

import (
	"testing"

	app "github.com/libmonsoon-dev/scanerr"
	"github.com/libmonsoon-dev/scanerr/config"
)

func Benchmark(b *testing.B) {
	conf := config.DefaultConfig()
	conf.PackagesLoaderConf.Patterns = []string{"./testdata/file-not-found/cmd"}
	s := app.NewScanerr(conf)

	inputErr := "runtime error: open /not-exist: open /not-exist: file does not exist"
	_, err := s.Scan(inputErr)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := s.Scan(inputErr)
		if err != nil {
			b.Fatal(err)
		}
	}
}
