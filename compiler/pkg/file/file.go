package file

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"

	"github.com/Akatsuki-py/gbdk-go/compiler/pkg/block"
	"github.com/Akatsuki-py/gbdk-go/compiler/pkg/function"
	"github.com/Akatsuki-py/gbdk-go/compiler/pkg/util"

	"golang.org/x/tools/go/analysis"
)

func CompileFile(f *ast.File, pass *analysis.Pass) (string, string, error) {
	result := ""
	decl := ""
	errMessage := ""

	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.ImportSpec:
			name := getModulePath(n)
			if !isGBDK(name) && !isAssetPackage(name) {
				errMessage = InvalidModule
			}
			result += include(name)
			decl += include(name)
			return false
		case *ast.FuncDecl:
			fun := function.CompileFunc(n, pass)
			result += fun.String()
			return false
		case *ast.ValueSpec:
			value, d := block.CompileValueSpec(n, pass, true)
			result += value + ";\n"
			decl += d + ";\n"
			return false
		}
		return true
	})

	if errMessage != "" {
		return "", "", errors.New(errMessage)
	}

	return result, decl, nil
}

func getModulePath(n *ast.ImportSpec) string {
	val := n.Path.Value
	return strings.ReplaceAll(val, "\"", "")
}

func isGBDK(name string) bool {
	for _, p := range util.GBDKPackage {
		if name == p {
			return true
		}
	}
	return false
}

func isAssetPackage(name string) bool {
	s := strings.Split(name, "/")
	if len(s) == 0 {
		return false
	}
	return s[len(s)-1] == "asset"
}

func include(name string) string {
	list := []string{
		"macro", "mem",
	}
	s := strings.Split(name, "/")
	h := s[len(s)-1]
	switch {
	case h == "gb":
		return fmt.Sprintf("#include <gb/%s.h>\n", h)
	case h == "cgb":
		return fmt.Sprintf("#include <gb/%s.h>\n", h)
	case h == "drawing":
		return "#include <gb/drawing.h>\n"
	case h == "str":
		return "#include <string.h>\n"
	case h == "asset":
		return "#include \"asset.h\"\n"
	case util.Contains(list, h):
		return ""
	default:
		return fmt.Sprintf("#include <%s.h>\n", h)
	}
}
