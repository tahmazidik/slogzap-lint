package analyzer

import "go/ast"

var logLevels = map[string]bool{
	"Debug": true,
	"Info":  true,
	"Warn":  true,
	"Error": true,
}

func isSlogCall(call *ast.CallExpr) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}

	xIdent, ok := sel.X.(*ast.Ident)
	if !ok || xIdent.Name != "slog" {
		return "", false
	}

	if !logLevels[sel.Sel.Name] {
		return "", false
	}

	return sel.Sel.Name, true
}

func isZapCall(call *ast.CallExpr, zapPresent bool) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}
	if !logLevels[sel.Sel.Name] {
		return "", false
	}

	if !zapPresent && !exprLooksLikeZap(sel.X) {
		return "", false
	}

	return sel.Sel.Name, true
}
