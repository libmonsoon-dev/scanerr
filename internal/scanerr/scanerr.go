package scanerr

import (
	"fmt"

	"github.com/libmonsoon-dev/scanerr/config"
	"github.com/libmonsoon-dev/scanerr/internal/source/packages"
	v1 "github.com/libmonsoon-dev/scanerr/model/v1"
)

func NewScanner(packagesLoader PackagesLoader, stringsExtractor SourceStringsExtractor,
	stringMatcher SourceStringMatcher) *Scanner {
	s := &Scanner{
		packagesLoader:   packagesLoader,
		stringsExtractor: stringsExtractor,
		stringMatcher:    stringMatcher,
	}

	return s
}

type Scanner struct {
	packagesLoader   PackagesLoader
	stringsExtractor SourceStringsExtractor
	stringMatcher    SourceStringMatcher
}

func (s *Scanner) Scan(inputErr string, config config.ScannerConfig) (*v1.Result, error) {
	pkgs, err := s.packagesLoader.Load(config.PackagesLoaderConf)
	if err != nil {
		return nil, fmt.Errorf("load packages: %w", err)
	}
	fmt.Printf("matched %v package(s) (%v total loaded)\n", len(pkgs), packages.Count(pkgs))

	strings := s.stringsExtractor.ExtractStrings(pkgs)
	fmt.Printf("extracted %d strings\n", len(strings))

	var bounds [][2]int
	strings, bounds = s.stringMatcher.FilterMatched(inputErr, strings)
	fmt.Printf("matched %d strings\n", len(strings))

	result := &v1.Result{
		InputErr: inputErr,
		List:     findUsages(strings, bounds),
	}

	return result, nil
}
