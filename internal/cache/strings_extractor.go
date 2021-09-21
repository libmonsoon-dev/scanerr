package cache

import (
	"sync"

	"golang.org/x/tools/go/packages"

	"github.com/libmonsoon-dev/scanerr/internal/source/ast"
)

type UpstreamStringExtractorFactory interface {
	CreateForPackage(p *packages.Package) *ast.StringExtractor
}

func NewStringExtractorFactory(upstream UpstreamStringExtractorFactory) *StringExtractorFactory {
	return &StringExtractorFactory{
		upstream: upstream,
		cache:    make(map[*packages.Package]*ast.StringExtractor),
	}
}

type StringExtractorFactory struct {
	upstream UpstreamStringExtractorFactory

	cacheMutex sync.Mutex
	cache      map[*packages.Package]*ast.StringExtractor
}

func (f *StringExtractorFactory) CreateForPackage(p *packages.Package) *ast.StringExtractor {
	f.cacheMutex.Lock()
	defer f.cacheMutex.Unlock()

	if extractor, ok := f.cache[p]; ok {
		return extractor
	}

	extractor := f.upstream.CreateForPackage(p)
	f.cache[p] = extractor

	return extractor
}

var _ UpstreamStringExtractorFactory = (*StringExtractorFactory)(nil)
