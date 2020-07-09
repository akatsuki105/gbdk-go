package expr

import (
	"fmt"
	"gbdk/compiler/pkg/types"
	"gbdk/compiler/pkg/util"
	"go/ast"
)

func CompileCallExpr(call *ast.CallExpr) string {
	fun := call.Fun
	funcName := CompileExpr(fun)
	if util.IsMacroStmt(funcName) {
		return funcName
	}
	if util.IsMemFunc(funcName) {
		return compileMemCall(call)
	}

	args := compileCallArgs(call.Args)

	name := util.RemoveTypePackage(funcName)
	if types.IsTypeFunc(name) {
		return fmt.Sprintf("(%s)(%s)", name, args)
	}
	return fmt.Sprintf("%s(%s)", funcName, args)
}

func CompileSelectExpr(selExpr *ast.SelectorExpr) string {
	sel := selExpr.Sel.Name
	return util.GetCFunc(sel)
}

func compileCallArgs(args []ast.Expr) string {
	result := ""
	for i, expr := range args {
		result += CompileExpr(expr)
		if i < len(args)-1 {
			result += ", "
		}
	}

	return result
}
