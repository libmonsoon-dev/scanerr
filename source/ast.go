package source

import (
	"fmt"
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"

	"github.com/libmonsoon-dev/scanerr/astutils"
)

func newAstStringExtractor(pkg *packages.Package) *astStringExtractor {
	v := &astStringExtractor{
		inspector:  inspector.New(pkg.Syntax),
		currentPkg: pkg,
	}
	return v
}

type astStringExtractor struct {
	result     []String
	inspector  *inspector.Inspector
	currentPkg *packages.Package
}

func (v *astStringExtractor) Extract() *astStringExtractor {
	types := []ast.Node{(*ast.BasicLit)(nil)}
	v.inspector.WithStack(types, v.visit)
	return v
}

func (v *astStringExtractor) visit(node ast.Node, push bool, stack []ast.Node) (proceed bool) {
	if !push || astutils.IsInImport(stack) {
		return
	}
	proceed = true

	basicLit, isBasicLit := astutils.AsBasicLit(node)
	if !isBasicLit || !astutils.IsStringLit(basicLit) {
		return
	}

	str, err := strconv.Unquote(basicLit.Value)
	if err != nil {
		fmt.Printf("unquote value %s: %v\n", basicLit.Value, err)
		return
	}

	if str == "" {
		return
	}

	v.addString(str, basicLit, stack)
	return
}

func (v *astStringExtractor) addString(str string, node *ast.BasicLit, stack []ast.Node) {
	result := String{
		Value:   str,
		Node:    node,
		Stack:   stack,
		Package: v.currentPkg,
	}

	v.result = append(v.result, result)
}

func (v *astStringExtractor) Result() []String {
	return v.result
}
