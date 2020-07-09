package main

import (
	"bytes"
	"gbdk/compiler/pkg/file"
	"go/ast"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gbdkgo",
	Doc:      `compile golang into gbdk's C`,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func main() {
	singlechecker.Main(
		Analyzer,
	)
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	isAsset, _ := strconv.ParseBool(os.Getenv("ASSET"))
	header := ""

	clangs := map[string]string{}
	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			name := fileName(pass, n)
			clang, decl, err := file.CompileFile(n, pass)
			if err != nil {
				panic(err)
			}
			header += decl
			clangs[name] = clang
		}
	})
	os.Mkdir("tmp", os.ModePerm)
	for key, value := range clangs {
		ioutil.WriteFile("tmp/"+key, []byte(value), 0666)
	}
	if isAsset {
		ioutil.WriteFile("tmp/asset.h", []byte(header), 0666)
	}

	return nil, nil
}

func fileName(pass *analysis.Pass, file *ast.File) string {
	var buf bytes.Buffer
	ast.Fprint(&buf, pass.Fset, file.Name.NamePos, nil)
	pos := buf.String()
	path := strings.Split(pos, ":")[0]
	paths := strings.Split(path, "/")
	fileName := strings.ReplaceAll(paths[len(paths)-1], ".go", ".c")
	return fileName
}
