package rules

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

type Rule interface {
	Check(call *ast.CallExpr, msg string) *analysis.Diagnostic
}
