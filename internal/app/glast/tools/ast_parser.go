package tools

import (
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/wispeeer/glast/internal/app/glast/utils"
)

var (
	fileset     *token.FileSet      = token.NewFileSet()
	packageDeps map[string][]string = make(map[string][]string)
)

func ASTParser() {
	for _, file := range GoFile {
		pkg := filepath.Dir(file)
		f, err := parser.ParseFile(fileset, file, nil, 0)
		if err != nil {
			panic(err)
		}
		var importPkg []string
		for _, i := range f.Imports {
			pkgName := utils.StrClean(i.Path.Value)
			importPkg = append(importPkg, pkgName)
		}
		packageDeps[pkg] = append(packageDeps[pkg], importPkg...)
	}
}
