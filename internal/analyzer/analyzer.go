package analyzer

import (
	"go/ast"

	"example.com/slogzaplint/internal/rules"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "logmsg",
	Doc:  "enforces log message format for slog and zap",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		zapPresent := fileImportsZap(file)

		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// slog
			if level, ok := isSlogCall(call); ok {
				checkAndReport(pass, call, "slog", level)
				return true
			}

			// zap
			if level, ok := isZapCall(call, zapPresent); ok {
				checkAndReport(pass, call, "zap", level)
				return true
			}

			return true
		})
	}

	return nil, nil
}

func checkAndReport(pass *analysis.Pass, call *ast.CallExpr, prefix, level string) {
	if len(call.Args) == 0 {
		return
	}
	msg, ok := stringLiteral(call.Args[0])
	if !ok {
		return
	}
	reportViolations(pass, call, prefix, level, rules.ValidateMessage(msg))
}
