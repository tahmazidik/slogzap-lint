package analyzer

import (
	"go/ast"
	"strconv"
)

func fileImportsZap(file *ast.File) bool {
	for _, imp := range file.Imports {
		path, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			continue
		}
		if path == "go.uber.org/zap" {
			return true
		}
	}
	return false
}

func exprLooksLikeZap(expr ast.Expr) bool {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return false
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	ident, ok := sel.X.(*ast.Ident)
	return ok && ident.Name == "zap"
}
