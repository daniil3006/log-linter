package rules

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

type LatinRule struct {
}

func NewLatinRule() *LatinRule {
	return &LatinRule{}
}

func (rule *LatinRule) Check(call *ast.CallExpr, s string) *analysis.Diagnostic {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.In(r, unicode.Latin) {
			return &analysis.Diagnostic{
				Pos:     call.Pos(),
				Message: "log message must be in english",
			}
		}
	}
	return nil
}
