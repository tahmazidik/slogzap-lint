package gclplugin

import (
	"fmt"

	"example.com/slogzaplint/internal/analyzer"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

type Plugin struct {
	cfg analyzer.Settings
}

func New(settings any) (register.LinterPlugin, error) {
	cfg := analyzer.DefaultSettings()

	if settings != nil {
		decoded, err := decodeSettings(settings, cfg)
		if err != nil {
			return nil, err
		}
		cfg = decoded
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid logmsg settings: %w", err)
	}

	return &Plugin{cfg: cfg}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.New(p.cfg)}, nil
}

func (p *Plugin) GetLoadMode() string { return "types" }
