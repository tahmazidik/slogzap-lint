package analyzer

import (
	"go/ast"
	"strings"

	"example.com/slogzaplint/internal/rules"
	"golang.org/x/tools/go/analysis"
)

func reportViolations(pass *analysis.Pass, call *ast.CallExpr, prefix, level string, violations []rules.Violation) {
	lvl := strings.ToLower(level)
	pos := call.Lparen
	if len(call.Args) > 0 {
		pos = call.Args[0].Pos()
	}
	for _, v := range violations {
		pass.Reportf(pos, "%s %s: %s", prefix, lvl, v.Message)
	}
}
