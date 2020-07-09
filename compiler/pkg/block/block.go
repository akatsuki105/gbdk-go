package block

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/Akatsuki-py/gbdk-go/compiler/pkg/expr"
	"github.com/Akatsuki-py/gbdk-go/compiler/pkg/macro"
	"github.com/Akatsuki-py/gbdk-go/compiler/pkg/util"

	"golang.org/x/tools/go/analysis"
)

var Pass *analysis.Pass

func CompileBlock(block *ast.BlockStmt, pass *analysis.Pass) string {
	Pass = pass
	result := []string{}
	ast.Inspect(block, func(n ast.Node) bool {
		list := block.List
		for _, stmt := range list {
			result = append(result, CompileStmt(stmt, pass))
		}
		return false
	})

	body := ""
	for _, stmt := range result {
		body += stmt + ";\n"
	}
	return fmt.Sprintf(`{
%s}`, body)
}

func CompileStmt(n ast.Stmt, pass *analysis.Pass) string {
	switch stmt := n.(type) {
	case *ast.ExprStmt:
		return expr.CompileExpr(stmt.X)
	case *ast.AssignStmt:
		return compileAssign(stmt)
	case *ast.ReturnStmt:
		return CompileReturnStmt(stmt, pass)
	case *ast.ForStmt:
		return CompileForStmt(stmt, pass)
	case *ast.IncDecStmt:
		return CompileIncDecStmt(stmt, pass)
	case *ast.IfStmt:
		return CompileIfStmt(stmt, pass)
	case *ast.SwitchStmt:
		return CompileSwitchStmt(stmt, pass)
	case *ast.BlockStmt:
		return CompileBlock(stmt, pass)
	case *ast.CaseClause:
		return CompileCaseClause(stmt, pass)
	default:
		result := ""
		ast.Inspect(stmt, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.ValueSpec:
				value, _ := CompileValueSpec(n, pass, false)
				result += value + "\n"
				return false
			default:
				return true
			}
		})
		return result
	}
}

func CompileCaseClause(stmt *ast.CaseClause, pass *analysis.Pass) string {
	condC := ""
	if stmt.List != nil {
		condExpr := stmt.List[0]
		condC = expr.CompileExpr(condExpr)
	}

	bodyC := ""
	for _, s := range stmt.Body {
		bodyC += CompileStmt(s, pass) + ";\n"
	}
	bodyC += "break"

	if condC != "" {
		return fmt.Sprintf(`case %s:
%s`, condC, bodyC)
	}
	return fmt.Sprintf(`default:
%s`, bodyC)
}

func CompileCaseClauseIntoIf(stmt *ast.CaseClause, index int) string {
	condC := ""
	if stmt.List != nil {
		condExpr := stmt.List[0]
		condC = expr.CompileExpr(condExpr)
	}

	bodyC := ""
	for _, s := range stmt.Body {
		bodyC += CompileStmt(s, Pass) + ";\n"
	}

	if condC == "" {
		result := fmt.Sprintf(` else {
%s}`, bodyC)
		return result
	}
	if index == 0 {
		result := fmt.Sprintf(`if (%s) {
%s}`, condC, bodyC)
		return result
	}
	result := fmt.Sprintf(` else if (%s) {
%s}`, condC, bodyC)
	return result
}

func CompileSwitchStmt(stmt *ast.SwitchStmt, pass *analysis.Pass) string {
	Pass = pass

	if stmt.Tag == nil {
		result := ""
		for i, s := range stmt.Body.List {
			caseCaluse := s.(*ast.CaseClause)
			result += CompileCaseClauseIntoIf(caseCaluse, i)
		}
		return result
	}

	tagC := expr.CompileExpr(stmt.Tag)
	bodyC := CompileBlock(stmt.Body, pass)

	return fmt.Sprintf("switch (%s) %s", tagC, bodyC)
}

func CompileIfStmt(stmt *ast.IfStmt, pass *analysis.Pass) string {
	Pass = pass

	cond := stmt.Cond
	condC := ""
	if cond != nil {
		condC = expr.CompileExpr(cond)
	}

	body := stmt.Body
	bodyC := CompileBlock(body, pass)

	elseNode := stmt.Else
	elseC := ""
	if elseNode != nil {
		elseC = CompileStmt(elseNode, pass)
	}

	if elseC == "" {
		result := fmt.Sprintf(`if (%s) %s`, condC, bodyC)
		return result
	}

	result := fmt.Sprintf(`if (%s) %s else %s`, condC, bodyC, elseC)
	return result
}

func CompileIncDecStmt(stmt *ast.IncDecStmt, pass *analysis.Pass) string {
	Pass = pass
	x := stmt.X
	return expr.CompileExpr(x) + stmt.Tok.String()
}

func CompileForStmt(forStmt *ast.ForStmt, pass *analysis.Pass) string {
	Pass = pass

	initC := ""
	if forStmt.Init != nil {
		initStmt := forStmt.Init
		initC = CompileStmt(initStmt, pass)
	}

	condC := ""
	if forStmt.Cond != nil {
		condExpr := forStmt.Cond
		condC = expr.CompileExpr(condExpr)
	}

	postC := ""
	if forStmt.Post != nil {
		postStmt := forStmt.Post
		postC = CompileStmt(postStmt, pass)
	}

	bodyC := ""
	if forStmt.Body != nil {
		bodyStmt := forStmt.Body
		bodyC = CompileBlock(bodyStmt, pass)
	}

	result := fmt.Sprintf(`for (%s; %s; %s)
%s`, initC, condC, postC, bodyC)
	return result
}

func CompileReturnStmt(returnStmt *ast.ReturnStmt, pass *analysis.Pass) string {
	Pass = pass
	results := returnStmt.Results
	if len(results) == 0 {
		return "return"
	}

	returnValue := results[0]
	return "return " + expr.CompileExpr(returnValue)
}

// CompileValueSpec func
// go: var i gb.UINT8 = 0
// c: UINT8 i = 0;
func CompileValueSpec(valueSpec *ast.ValueSpec, pass *analysis.Pass, global bool) (string, string) {
	Pass = pass
	varName := valueSpec.Names[0].Name

	if valueSpec.Values == nil {
		varTypeExpr := valueSpec.Type
		varType := util.RemoveTypePackage(Pass.TypesInfo.TypeOf(varTypeExpr).String())

		result := fmt.Sprintf("%s %s", varType, varName)
		return result, ""
	}

	varValueExpr := valueSpec.Values[0]
	varType := util.RemoveTypePackage(Pass.TypesInfo.TypeOf(varValueExpr).String())
	isArray, tmp := util.IsArray(varType)
	if isArray {
		varName += "[]"
		varType = tmp
	}
	varValue := expr.CompileExpr(varValueExpr)

	// define stmt
	if strings.HasPrefix(varValue, "Define") {
		stmt := macro.CompileDefine(varName, varValueExpr)
		return stmt, stmt
	}

	// var stmt
	result := fmt.Sprintf("%s %s = %s", varType, varName, varValue)
	decl := fmt.Sprintf("extern %s %s", varType, varName)
	return result, decl
}

// go: i := 0
// c: int i = 0;
func compileAssign(assign *ast.AssignStmt) string {
	lhs := assign.Lhs[0]
	varName := expr.CompileExpr(lhs)

	rhs := assign.Rhs[0]
	varType := util.RemoveTypePackage(Pass.TypesInfo.TypeOf(rhs).String())
	if varType == "string" {
		varType = "char*"
	} else if cType, ok := util.CType[varType]; ok {
		varType = cType
	}
	varValue := ""
	for _, e := range assign.Rhs {
		varValue += expr.CompileExpr(e)
	}

	if assign.Tok.String() == ":=" {
		return fmt.Sprintf("%s %s = %v", varType, varName, varValue)
	}
	return fmt.Sprintf("%s %s %v", varName, assign.Tok.String(), varValue)
}
