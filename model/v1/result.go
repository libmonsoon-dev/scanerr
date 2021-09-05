package v1

import (
	"io"
	"strconv"
	"strings"

	"github.com/labstack/gommon/color"
)

type Result struct {
	InputErr string
	List     []ResultsListItem
}

func (r Result) String() string {
	var builder strings.Builder

	for _, item := range r.List {
		builder.WriteString(item.FuncName)
		builder.WriteRune('\t')
		highlightError(r.InputErr, item.MatchBounds[0], item.MatchBounds[1], &builder)
		builder.WriteRune('\n')

		builder.WriteRune('\t')
		builder.WriteString(item.Filename)
		builder.WriteRune(':')
		builder.WriteString(strconv.Itoa(item.Line))
		builder.WriteRune(':')
		builder.WriteString(strconv.Itoa(item.Column))
		builder.WriteRune('\n')
	}

	return builder.String()
}

type ResultsListItem struct {
	Filename    string
	Line        int
	Column      int
	FuncName    string
	MatchBounds [2]int
}

func highlightError(inputErr string, start, end int, out io.StringWriter) {
	ignoreWriteError(out.WriteString(inputErr[:start]))
	matchPath := color.Underline(inputErr[start:end])
	ignoreWriteError(out.WriteString(matchPath))
	ignoreWriteError(out.WriteString(inputErr[end:]))
}

func ignoreWriteError(_ int, _ error) {}
