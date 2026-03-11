package rules

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type SensitiveDataRule struct {
	sensitiveWords []string
}

func NewSensitiveDataRule(sensitiveWords []string) *SensitiveDataRule {
	rule := &SensitiveDataRule{
		sensitiveWords: []string{
			"password",
			"api_key",
			"token",
		},
	}
	if len(sensitiveWords) == 0 {
		return rule
	}

	for i := range sensitiveWords {
		sensitiveWords[i] = strings.ToLower(sensitiveWords[i])
	}

	rule.sensitiveWords = append(rule.sensitiveWords, sensitiveWords...)
	return rule
}

func (rule *SensitiveDataRule) Check(call *ast.CallExpr, s string) *analysis.Diagnostic {
	if rule.isContainsSensitiveWords(s) && rule.hasVariables(call) {
		return &analysis.Diagnostic{
			Pos:     call.Pos(),
			Message: "log message must not contain sensitive data",
		}
	}
	return nil
}

func (rule *SensitiveDataRule) isContainsSensitiveWords(s string) bool {
	s = strings.ToLower(s)

	for _, word := range rule.sensitiveWords {
		if strings.Contains(s, word) {
			return true
		}
	}
	return false
}

func (rule *SensitiveDataRule) hasVariables(call *ast.CallExpr) bool {
	for _, arg := range call.Args {
		switch arg.(type) {
		case *ast.Ident:
			return true
		case *ast.BinaryExpr:
			return true
		case *ast.SelectorExpr:
			return true
		case *ast.CallExpr:
			return true
		}
	}
	return false
}
