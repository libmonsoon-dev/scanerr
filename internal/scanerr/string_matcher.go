package scanerr

import "github.com/libmonsoon-dev/scanerr/internal/source"

type SourceStringMatcher interface {
	FilterMatched(originalError string, input []source.String) (filtered []source.String, matches [][2]int)
}
