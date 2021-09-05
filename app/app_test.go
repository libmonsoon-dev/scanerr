package app_test

import (
	"testing"

	"github.com/libmonsoon-dev/scanerr/app"
	"github.com/libmonsoon-dev/scanerr/config"
)

func TestScanerr(t *testing.T) {
	conf := config.DefaultConfig()
	conf.PackagesLoaderConf.Patterns = []string{"../testdata/file-not-found/cmd"}
	s := app.NewScanerr(conf)

	inputErr := "runtime error: open /not-exist: open /not-exist: file does not exist"
	result, err := s.Scan(inputErr)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
