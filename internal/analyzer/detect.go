package analyzer

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

func isSlogCall(pass *analysis.Pass, call *ast.CallExpr, levels map[string]bool) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}

	level := sel.Sel.Name
	if !levels[level] {
		return "", false
	}

	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return "", false
	}

	obj := pass.TypesInfo.Uses[id]
	pkgName, ok := obj.(*types.PkgName)
	if !ok {
		return "", false
	}

	if pkgName.Imported().Path() != "log/slog" {
		return "", false
	}

	return level, true
}

func isZapCall(pass *analysis.Pass, call *ast.CallExpr, levels map[string]bool) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}

	level := sel.Sel.Name
	if !levels[level] {
		return "", false
	}

	selection := pass.TypesInfo.Selections[sel]
	if selection == nil {
		return "", false
	}

	if selection.Kind() != types.MethodVal {
		return "", false
	}

	recv := selection.Recv()
	if recv == nil {
		return "", false
	}

	t := recv
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	named, ok := t.(*types.Named)
	if !ok {
		return "", false
	}

	obj := named.Obj()
	pkg := obj.Pkg()
	if pkg == nil {
		return "", false
	}

	if pkg.Path() != "go.uber.org/zap" {
		return "", false
	}

	switch obj.Name() {
	case "Logger", "SugaredLogger":
		return level, true
	default:
		return "", false
	}
}
