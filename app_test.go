package scanerr_test

import (
	"testing"

	"go.uber.org/goleak"

	app "github.com/libmonsoon-dev/scanerr"
	"github.com/libmonsoon-dev/scanerr/config"
)

func TestScanerr(t *testing.T) {
	defer goleak.VerifyNone(t)

	conf := config.DefaultConfig()
	conf.PackagesLoaderConf.Patterns = []string{"./testdata/file-not-found/cmd"}
	s := app.NewScanerr(conf)

	inputErr := "runtime error: open /not-exist: open /not-exist: file does not exist"
	result, err := s.Scan(inputErr)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
