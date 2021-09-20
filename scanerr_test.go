package scanerr_test

import (
	"testing"

	"go.uber.org/goleak"

	app "github.com/libmonsoon-dev/scanerr"
	"github.com/libmonsoon-dev/scanerr/config"
)

func TestScanerr(t *testing.T) {
	defer goleak.VerifyNone(t)

	appConf := config.DefaultAppConfig()
	s := app.NewScanerr(appConf)

	scannerConf := config.DefaultScannerConfig()
	scannerConf.PackagesLoaderConf.Patterns = []string{"./testdata/file-not-found/cmd"}

	inputErr := "runtime error: open /not-exist: open /not-exist: file does not exist"
	result, err := s.Scan(inputErr, scannerConf)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
