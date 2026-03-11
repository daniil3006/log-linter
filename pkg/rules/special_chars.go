package rules

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

type SpecialCharsRule struct {
	allowedCharacters map[rune]bool
}

func NewSpecialCharsRule() *SpecialCharsRule {
	return &SpecialCharsRule{
		allowedCharacters: map[rune]bool{
			'%': true,
			'_': true,
		},
	}
}

func (rule *SpecialCharsRule) Check(call *ast.CallExpr, s string) *analysis.Diagnostic {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) && !rule.allowedCharacters[r] {
			return &analysis.Diagnostic{
				Pos:     call.Pos(),
				Message: "log message must not contain special characters",
			}
		}
	}
	return nil
}
