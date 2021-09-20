package source

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/packages"
)

type String struct {
	Value   string
	Node    *ast.BasicLit
	Package *packages.Package
	Stack   []ast.Node
}

func (s String) Position() token.Position {
	return s.Package.Fset.Position(s.Node.Pos())
}
