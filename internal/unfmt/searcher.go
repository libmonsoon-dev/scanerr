package unfmt

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/libmonsoon-dev/scanerr/internal/source"
)

func NewStringsMatcher() *StringsMatcher {
	s := &StringsMatcher{}

	return s
}

type StringsMatcher struct {
}

func (m *StringsMatcher) FilterMatched(originalError string, input []source.String) ([]source.String, [][2]int) {
	regexpList := make([]*regexp.Regexp, len(input))
	var err error
	for i := range input {
		regexpExpr := m.unfmt(input[i].Value)

		regexpList[i], err = regexp.Compile(regexpExpr)
		if err != nil {
			fmt.Printf("compile %q regexp: %s\n", regexpExpr, err)
		}
	}

	var matches [][2]int
	// TODO: change input type from slice to channel
	input, matches = filter(originalError, input, regexpList)
	sortStrings(input, matches)

	return input, matches

}

type byMatchIndex struct {
	strings []source.String
	matches [][2]int
}

func (s byMatchIndex) Len() int {
	return len(s.strings)
}

func (s byMatchIndex) Less(i, j int) bool {
	if s.matches[i][0] != s.matches[j][0] {
		return s.matches[i][0] < s.matches[j][0]
	}

	return s.matches[i][1] > s.matches[j][1]
}

func (s byMatchIndex) Swap(i, j int) {
	s.strings[i], s.strings[j] = s.strings[j], s.strings[i]
	s.matches[i], s.matches[j] = s.matches[j], s.matches[i]
}

func sortStrings(input []source.String, matches [][2]int) {
	if len(input) != len(matches) {
		panic("Invalid input length")
	}

	sort.Stable(byMatchIndex{input, matches})
}

func (m *StringsMatcher) unfmt(value string) string {
	value = regexp.QuoteMeta(value)

	// TODO: unfmt
	value = strings.ReplaceAll(value, "%v", ".*")
	value = strings.ReplaceAll(value, "%w", ".*")

	return value
}

func filter(originalError string, input []source.String, regexpList []*regexp.Regexp) ([]source.String, [][2]int) {
	n := 0

	matches := make([][2]int, 0)

	var match []int
	for i := range regexpList {
		if regexpList[i] == nil {
			continue
		}

		match = regexpList[i].FindStringIndex(originalError)
		if match == nil || len(match) != 2 {
			continue
		}

		input[n] = input[i]
		regexpList[n] = regexpList[i]
		matches = append(matches, *(*[2]int)(match))
		n++
	}

	return input[:n], matches
}
