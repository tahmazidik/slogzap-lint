package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func reportViolations(pass *analysis.Pass, call *ast.CallExpr, prefix, level string, violations []string) {
	lvl := strings.ToLower(level)
	for _, v := range violations {
		pass.Reportf(call.Lparen, "%s %s: %s", prefix, lvl, v)
	}
}
