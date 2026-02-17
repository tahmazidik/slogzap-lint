package analyzer

import (
	"go/ast"

	"example.com/slogzaplint/internal/rules"
	"golang.org/x/tools/go/analysis"
)

func New(cfg Settings) *analysis.Analyzer {
	levelSet := make(map[string]bool, len(cfg.Levels))
	for _, l := range cfg.Levels {
		levelSet[l] = true
	}

	return &analysis.Analyzer{
		Name: "logmsg",
		Doc:  "enforces log message format for slog and zap",
		Run: func(pass *analysis.Pass) (any, error) {
			//Чтобы неверны конфиг сразу выдавал ошибку
			if err := cfg.Validate(); err != nil {
				return nil, err
			}

			for _, file := range pass.Files {
				ast.Inspect(file, func(n ast.Node) bool {
					call, ok := n.(*ast.CallExpr)
					if !ok {
						return true
					}

					if level, ok := isSlogCall(pass, call, levelSet); ok {
						checkAndReport(pass, call, "slog", level, cfg)
						return true
					}
					if level, ok := isZapCall(pass, call, levelSet); ok {
						checkAndReport(pass, call, "zap", level, cfg)
						return true
					}
					return true
				})
			}
			return nil, nil
		},
	}
}

func checkAndReport(pass *analysis.Pass, call *ast.CallExpr, prefix, level string, cfg Settings) {
	if len(call.Args) == 0 {
		return
	}
	msg, ok := stringLiteral(call.Args[0])
	if !ok {
		return
	}

	violations := rules.ValidateMessageWithSensitiveKeys(msg, cfg.SensitiveKeys)

	reportViolations(pass, call, prefix, level, violations)
}
