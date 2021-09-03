package source

import (
	"go/ast"

	"golang.org/x/tools/go/packages"
)

type String struct {
	Value   string
	Node    *ast.BasicLit
	Package *packages.Package
	Stack   []ast.Node
}
