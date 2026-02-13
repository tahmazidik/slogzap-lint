package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "analyzer",
	Doc:  "checks log messages for slog and zap (MVP: just directs calls)",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			pass.Reportf(call.Lparen, "found a call expression")
			return true
		})
	}
	return nil, nil
}
