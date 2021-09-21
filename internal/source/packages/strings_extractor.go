package packages

import (
	"golang.org/x/tools/go/packages"

	"github.com/libmonsoon-dev/scanerr/config"
	"github.com/libmonsoon-dev/scanerr/internal/semaphore"
	"github.com/libmonsoon-dev/scanerr/internal/source"
	"github.com/libmonsoon-dev/scanerr/internal/source/ast"
)

type ASTStringExtractorFactory interface {
	CreateForPackage(p *packages.Package) *ast.StringExtractor
}

func NewStringsExtractor(factory ASTStringExtractorFactory) *StringsExtractor {
	s := &StringsExtractor{
		factory: factory,
	}

	return s
}

type StringsExtractor struct {
	factory ASTStringExtractorFactory
}

func (s *StringsExtractor) ExtractStrings(pks []*packages.Package) (result []source.String) {
	packages.Visit(pks, nil, func(p *packages.Package) {
		result = append(result, s.extractStrings(p)...)
	})

	return
}

func (s *StringsExtractor) extractStrings(p *packages.Package) []source.String {
	return s.factory.
		CreateForPackage(p).
		Extract()
}

func NewConcurrentStringsExtractor(conf config.StringsExtractorConfig, factory ASTStringExtractorFactory) *ConcurrentStringsExtractor {
	s := &ConcurrentStringsExtractor{
		conf:    conf,
		factory: factory,
	}

	return s
}

type ConcurrentStringsExtractor struct {
	conf    config.StringsExtractorConfig
	factory ASTStringExtractorFactory
}

func (s *ConcurrentStringsExtractor) ExtractStrings(pks []*packages.Package) (result []source.String) {
	resultCh := make(chan []source.String)
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

func (s *ConcurrentStringsExtractor) extractStrings(p *packages.Package) []source.String {
	return s.factory.
		CreateForPackage(p).
		Extract()
}
