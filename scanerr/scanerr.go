package scanerr

import (
	"fmt"

	"github.com/libmonsoon-dev/scanerr/packages"
	"github.com/libmonsoon-dev/scanerr/searcher"
	"github.com/libmonsoon-dev/scanerr/source"
)

func NewScanerr(conf Config, packagesLoader packages.Loader, stringsExtractor source.StringsExtractor,
	searcher searcher.Searcher) *Scanerr {
	s := &Scanerr{
		conf:             conf,
		packagesLoader:   packagesLoader,
		stringsExtractor: stringsExtractor,
		searcher:         searcher,
	}

	return s
}

type Scanerr struct {
	conf             Config
	packagesLoader   packages.Loader
	stringsExtractor source.StringsExtractor
	searcher         searcher.Searcher
}

func (s *Scanerr) Scan(inputErr string) ([]searcher.Result, error) {
	pkgs, err := s.packagesLoader.Load()
	if err != nil {
		return nil, fmt.Errorf("load packages: %w", err)
	}

	strings := s.stringsExtractor.ExtractStrings(pkgs)
	fmt.Printf("extracted %d strings\n", len(strings))

	result := s.searcher.Search(inputErr, strings)
	return result, nil
}
