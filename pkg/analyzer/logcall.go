package analyzer

import (
	"go/ast"
	"go/types"
	"log-linter/pkg/config"
	"log-linter/pkg/rules"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func isLogCall(pass *analysis.Pass, selector *ast.SelectorExpr) bool {
	obj, ok := pass.TypesInfo.Uses[selector.Sel]
	if !ok {
		return false
	}

	fn, ok := obj.(*types.Func)
	if !ok {
		return false
	}

	if fn.Pkg() == nil {
		return false
	}
	pkg := fn.Pkg().Path()
	pkgMethods, ok := methods[pkg]
	if !ok {
		return false
	}

	return pkgMethods[selector.Sel.Name]
}

func extractMsg(expr ast.Expr) (string, bool) {
	switch arg := expr.(type) {
	case *ast.BasicLit:
		str, err := strconv.Unquote(arg.Value)
		if err != nil {
			return "", false
		}
		return str, true
	case *ast.BinaryExpr:
		strX, okX := extractMsg(arg.X)
		strY, okY := extractMsg(arg.Y)

		if okX && okY {
			return strX + strY, true
		}

		if okX {
			return strX, true
		}

		if okY {
			return strY, true
		}

		return "", false
	case *ast.CallExpr:
		if len(arg.Args) > 0 {
			return extractMsg(arg.Args[0])
		}
	}

	return "", false
}

func fillActiveRules(cfg *config.Config) {
	if cfg.Lowercase {
		activeRules = append(activeRules, rules.NewLowercaseRule())
	}

	if cfg.Latin {
		activeRules = append(activeRules, rules.NewLatinRule())
	}

	if cfg.SpecialChars {
		activeRules = append(activeRules, rules.NewSpecialCharsRule())
	}

	if cfg.SensitiveData {
		activeRules = append(activeRules, rules.NewSensitiveDataRule(cfg.SensitiveWords))
	}
}

func messageIndex(method string) int {
	if strings.HasSuffix(method, "Context") {
		return 1
	}
	return 0
}
