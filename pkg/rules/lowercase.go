package rules

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

type LowercaseRule struct {
}

func NewLowercaseRule() *LowercaseRule {
	return &LowercaseRule{}
}

func (r *LowercaseRule) Check(call *ast.CallExpr, s string) *analysis.Diagnostic {
	if s == "" {
		return nil
	}

	firstLetter := rune(s[0])
	if firstLetter > unicode.MaxASCII {
		return nil
	}

	if unicode.IsLetter(firstLetter) && !unicode.IsLower(firstLetter) {
		fixed := strings.ToLower(string(firstLetter)) + s[1:]
		return &analysis.Diagnostic{
			Pos:     call.Args[0].Pos() + 1,
			Message: "log message must be in lowercase",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "convert to lowercase",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     call.Args[0].Pos(),
							End:     call.Args[0].End(),
							NewText: []byte(`"` + fixed + `"`),
						},
					},
				},
			},
		}
	}
	return nil
}
