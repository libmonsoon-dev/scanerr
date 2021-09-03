package searcher

import (
	"github.com/libmonsoon-dev/scanerr/source"
)

type Searcher interface {
	Search(inputErr string, strings []source.String) []Result
}

func NewSearcher() Searcher {
	s := &searcher{}

	return s
}

type searcher struct {
}

func (s *searcher) Search(inputErr string, strings []source.String) (result []Result) {

	return
}
