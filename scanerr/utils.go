package scanerr

import (
	"fmt"
	"go/ast"
	"go/token"

	v1 "github.com/libmonsoon-dev/scanerr/model/v1"
	"github.com/libmonsoon-dev/scanerr/source"
)

func findUsages(strings []source.String, bounds [][2]int) []v1.ResultsListItem { // TODO
	results := make([]v1.ResultsListItem, len(strings))
	var pos token.Position
	for i := range strings {
		pos = strings[i].Position()
		results[i] = v1.ResultsListItem{
			Filename:    pos.Filename,
			Line:        pos.Line,
			Column:      pos.Column,
			FuncName:    getFuncName(strings[i].Stack),
			MatchBounds: bounds[i],
		}
	}

	return results
}

func getFuncName(stack []ast.Node) string {
	for i := len(stack) - 1; i >= 0; i-- {
		fn, ok := stack[i].(*ast.FuncDecl)
		if !ok || fn.Name == nil {
			continue
		}
		return fmt.Sprintf("%s()", fn.Name.Name)
	}

	return "<package level>"
}
