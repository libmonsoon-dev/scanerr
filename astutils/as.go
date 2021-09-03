package astutils

import "go/ast"

func AsBasicLit(node ast.Node) (lit *ast.BasicLit, ok bool) {
	lit, ok = node.(*ast.BasicLit)
	return
}
