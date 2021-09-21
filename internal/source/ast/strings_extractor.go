package ast

import (
	"fmt"
	"go/ast"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"

	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"

	"github.com/libmonsoon-dev/scanerr/internal/astutils"
	"github.com/libmonsoon-dev/scanerr/internal/source"
)

func NewStringExtractorFactory() *StringExtractorFactory {
	return &StringExtractorFactory{}
}

type StringExtractorFactory struct{}

func (StringExtractorFactory) CreateForPackage(pkg *packages.Package) *StringExtractor {
	v := &StringExtractor{
		inspector:  inspector.New(pkg.Syntax),
		currentPkg: pkg,
	}
	return v
}

type StringExtractor struct {
	inspector  *inspector.Inspector
	currentPkg *packages.Package

	once   sync.Once
	result []source.String
}

func (v *StringExtractor) Extract() []source.String {
	v.once.Do(func() {
		types := []ast.Node{(*ast.BasicLit)(nil)}
		v.inspector.WithStack(types, v.visit)
	})
	return v.result
}

func (v *StringExtractor) visit(node ast.Node, push bool, stack []ast.Node) (proceed bool) {
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

	str = strings.TrimSpace(str)
	const minStrSize = 3
	if len(str) < minStrSize || !utf8.ValidString(str) {
		return
	}

	v.addString(str, basicLit, stack)
	return
}

func (v *StringExtractor) addString(str string, node *ast.BasicLit, stack []ast.Node) {
	result := source.String{
		Value:   str,
		Node:    node,
		Stack:   stack,
		Package: v.currentPkg,
	}

	v.result = append(v.result, result)
}
