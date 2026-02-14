package analyzer

import "go/ast"

func isSlogCall(call *ast.CallExpr, levels map[string]bool) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}

	xIdent, ok := sel.X.(*ast.Ident)
	if !ok || xIdent.Name != "slog" {
		return "", false
	}

	if !levels[sel.Sel.Name] {
		return "", false
	}

	return sel.Sel.Name, true
}

func isZapCall(call *ast.CallExpr, zapPresent bool, levels map[string]bool) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}

	if !levels[sel.Sel.Name] {
		return "", false
	}

	if !zapPresent && !exprLooksLikeZap(sel.X) {
		return "", false
	}

	return sel.Sel.Name, true
}
