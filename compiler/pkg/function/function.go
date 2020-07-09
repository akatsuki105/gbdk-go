package function

import (
	"fmt"
	"go/ast"

	"github.com/Akatsuki-py/gbdk-go/compiler/pkg/block"
	"github.com/Akatsuki-py/gbdk-go/compiler/pkg/expr"

	"golang.org/x/tools/go/analysis"
)

var Pass *analysis.Pass

type Function struct {
	Name   string
	Body   string
	Args   []Arg
	Return string
}

type Arg struct {
	Name string
	Type string
}

func (f *Function) String() string {
	args := ""
	for i, arg := range f.Args {
		if i == len(f.Args)-1 {
			args += fmt.Sprintf("%s %s", arg.Type, arg.Name)
		} else {
			args += fmt.Sprintf("%s %s, ", arg.Type, arg.Name)
		}
	}

	body := f.Body
	return fmt.Sprintf(`%s %s(%s) %s

`, f.Return, f.Name, args, body)
}

func CompileFunc(funcDecl *ast.FuncDecl, pass *analysis.Pass) Function {
	Pass = pass
	f := Function{
		Name: funcDecl.Name.String(),
		Body: "",
		Args: []Arg{},
	}

	ast.Inspect(funcDecl, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.FuncType:
			f.Args, f.Return = compileFuncType(n)
			return false
		case *ast.BlockStmt:
			f.Body = block.CompileBlock(n, pass)
			return false
		}
		return true
	})

	return f
}

func compileFuncType(typeDecl *ast.FuncType) ([]Arg, string) {
	params := typeDecl.Params
	args := compileParams(params)
	results := typeDecl.Results
	result := compileResult(results)
	return args, result
}

func compileParams(paramsDecl *ast.FieldList) []Arg {
	args := []Arg{}
	for _, field := range paramsDecl.List {
		for _, ident := range field.Names {
			a := Arg{
				Name: ident.Name,
			}
			a.Type = string(expr.CompileExpr(field.Type))
			args = append(args, a)
		}
	}

	return args
}

func compileResult(resultDecl *ast.FieldList) string {
	if resultDecl == nil {
		return "void"
	}

	for _, field := range resultDecl.List {
		return string(expr.CompileExpr(field.Type))
	}

	return string("void")
}
