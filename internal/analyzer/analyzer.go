package analyzer

import (
	"go/ast"

	"example.com/slogzaplint/internal/rules"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "analyzer",
	Doc:  "checks log messages for slog and zap (MVP: just directs calls)",
	Run:  run,
}

var slogMethods = map[string]bool{
	"Debug": true,
	"Info":  true,
	"Warn":  true,
	"Error": true,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Проверяем что вызывается селектор
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// X должен быть идентификатором slog
			xIdent, ok := sel.X.(*ast.Ident)
			if !ok || xIdent.Name != "slog" {
				return true
			}

			// sel должен быть одним из методов slog
			if !slogMethods[sel.Sel.Name] {
				return true
			}

			// 1-й аргумент должен суще и быть строковым литералом
			if len(call.Args) == 0 {
				return true
			}
			msg, ok := stringLiteral(call.Args[0])
			// тогда просто пропускаем не-литералы
			if !ok {
				return true
			}
			// репортим и валидируем
			violations := rules.ValidateMessage(msg)
			for _, v := range violations {
				pass.Reportf(call.Lparen, "slog message rule violation: %s", v)
			}
			return true
		})
	}
	return nil, nil
}
