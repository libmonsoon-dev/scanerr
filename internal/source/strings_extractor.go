package source

import (
	"golang.org/x/tools/go/packages"

	"github.com/libmonsoon-dev/scanerr/internal/semaphore"
)

type StringsExtractor interface {
	ExtractStrings([]*packages.Package) []String
}

func NewStringsExtractor(conf StringsExtractorConfig) StringsExtractor {
	s := &stringsExtractor{
		conf: conf,
	}

	return s
}

type StringsExtractorConfig struct {
	NumWorkers int
}

type stringsExtractor struct {
	conf StringsExtractorConfig
}

func (s *stringsExtractor) ExtractStrings(pks []*packages.Package) (result []String) {
	resultCh := make(chan []String)
	sem := semaphore.NewSemaphore(s.conf.NumWorkers)

	go func(pks []*packages.Package) {
		packages.Visit(pks, nil, func(p *packages.Package) {
			sem.Acquire()
			go func() {
				resultCh <- s.extractStrings(p)
				sem.Release()
			}()
		})

		// wait until all workers exited
		sem.AcquireAll()
		close(resultCh)
	}(pks)

	for msg := range resultCh {
		result = append(result, msg...)
	}

	return
}

func (s *stringsExtractor) extractStrings(p *packages.Package) []String {
	return newAstStringExtractor(p).
		Extract().
		Result()
}

var _ StringsExtractor = (*stringsExtractor)(nil)
