package expr

import (
	"fmt"
	"go/ast"
)

func compileMemCall(call *ast.CallExpr) string {
	fun := call.Fun
	funcName := CompileExpr(fun)
	switch funcName {
	case "ReadMemory":
		args := call.Args
		addr := CompileExpr(args[0])
		return fmt.Sprintf("*(unsigned char *)%s", addr)
	case "WriteMemory":
		args := call.Args
		addr := CompileExpr(args[0])
		value := CompileExpr(args[1])
		return fmt.Sprintf("*(unsigned char *)%s = %s", addr, value)
	default:
		return ""
	}
}
