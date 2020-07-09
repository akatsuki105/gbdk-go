package expr

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

var fset = token.NewFileSet()

func CompileExpr(expr ast.Expr) string {
	if expr == nil {
		return ""
	}

	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		return CompileBinaryExpr(expr)
	case *ast.ParenExpr:
		return CompileParenExpr(expr)
	case *ast.BasicLit:
		return expr.Value
	case *ast.Ident:
		return expr.Name
	case *ast.UnaryExpr:
		return CompileUnaryExpr(expr)
	case *ast.CallExpr:
		return CompileCallExpr(expr)
	case *ast.CompositeLit:
		return CompileCompositeLit(expr)
	case *ast.SelectorExpr:
		return CompileSelectExpr(expr)
	case *ast.IndexExpr:
		return CompileIndexExpr(expr)
	default:
		var buffer bytes.Buffer
		printer.Fprint(&buffer, fset, expr)
		return buffer.String()
	}
}

func CompileParenExpr(expr *ast.ParenExpr) string {
	return CompileExpr(expr.X)
}

func CompileIndexExpr(expr *ast.IndexExpr) string {
	ident := CompileExpr(expr.X)
	index := CompileExpr(expr.Index)
	return fmt.Sprintf("%s[%s]", ident, index)
}

func CompileUnaryExpr(expr *ast.UnaryExpr) string {
	var buffer bytes.Buffer
	printer.Fprint(&buffer, fset, expr)
	return buffer.String()
}

func CompileCompositeLit(expr *ast.CompositeLit) string {
	t := expr.Type
	_, ok := t.(*ast.ArrayType)
	if ok {
		if expr.Elts != nil {
			result := "{"
			for _, elt := range expr.Elts {
				result += CompileExpr(elt) + ","
			}
			result += "}"
			return result
		}
		var buffer bytes.Buffer
		printer.Fprint(&buffer, fset, expr)
		s := strings.Split(buffer.String(), "{")
		return "{" + s[len(s)-1]
	}
	return ""
}
