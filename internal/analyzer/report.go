package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func reportViolations(pass *analysis.Pass, call *ast.CallExpr, prefix, level string, violations []string) {
	lvl := strings.ToLower(level)
	for _, v := range violations {
		pos := call.Lparen
		if len(call.Args) > 0 {
			pos = call.Args[0].Pos()
		}
		pass.Reportf(pos, "%s %s: %s", prefix, lvl, v)
	}
}
