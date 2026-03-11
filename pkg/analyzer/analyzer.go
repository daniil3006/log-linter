package analyzer

import (
	"flag"
	"go/ast"
	"log-linter/pkg/config"
	"log-linter/pkg/rules"

	"golang.org/x/tools/go/analysis"
)

var (
	configPath string

	activeRules []rules.Rule

	Analyzer = &analysis.Analyzer{
		Name:  "loglinter",
		Doc:   "linter checks rules for log entries",
		Flags: flag.FlagSet{},
		Run:   run,
	}
)

func init() {
	Analyzer.Flags.StringVar(&configPath, "config", "loglinter.yaml", "path to config file")
}

func run(pass *analysis.Pass) (any, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, err
	}
	fillActiveRules(cfg)

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			selector, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			if isLogCall(pass, selector) {
				msgIdx := messageIndex(selector.Sel.Name)

				s, ok := extractMsg(call.Args[msgIdx])
				if !ok {
					return true
				}

				for _, r := range activeRules {
					diagnostic := r.Check(call, s)
					if diagnostic != nil {
						pass.Report(*diagnostic)
					}
				}
			}
			return true
		})
	}
	return nil, nil
}
