package astutils

import "go/ast"

const (
	KindString = "STRING"
)

func IsStringLit(blt *ast.BasicLit) bool {
	return blt.Kind.String() == KindString
}

func IsInImport(stack []ast.Node) bool {
	for _, node := range stack {
		if IsImportNode(node) {
			return true
		}
	}
	return false
}

func IsImportNode(node ast.Node) bool {
	_, ok := node.(*ast.ImportSpec)
	return ok
}
