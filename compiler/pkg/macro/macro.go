package macro

import (
	"fmt"
	"gbdk/compiler/pkg/expr"
	"go/ast"
)

func CompileDefine(name string, defineExpr ast.Expr) string {
	call := defineExpr.(*ast.CallExpr)
	args := call.Args
	if len(args) == 0 {
		return ""
	}
	value := expr.CompileExpr(args[0])
	return fmt.Sprintf("#define %s %s", name, value)
}
