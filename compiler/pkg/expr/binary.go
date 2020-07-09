package expr

import (
	"fmt"
	"go/ast"
	"strings"
)

func CompileBinaryExpr(expr *ast.BinaryExpr) string {
	x := CompileExpr(expr.X)
	y := CompileExpr(expr.Y)
	Op := expr.Op.String()

	xIsStr := strings.HasPrefix(x, "\"") && strings.HasSuffix(x, "\"")
	yIsStr := strings.HasPrefix(y, "\"") && strings.HasSuffix(y, "\"")
	if xIsStr && yIsStr {
		return strings.TrimRight(x, "\"") + strings.TrimLeft(y, "\"")
	}
	return fmt.Sprintf("%s %s %s", x, Op, y)
}
