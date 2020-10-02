package helper

import (
	"go/ast"

	"github.com/negrel/debuggo/pkg/inspector"
)

func ChangePackage(name string) inspector.Inspector {
	return func(node ast.Node) bool {
		pkg, isPackage := node.(*ast.Package)
		if !isPackage {
			return false
		}

		pkg.Name = name

		return false
	}
}

func ApplyOnTopDecl(inspectors ...inspector.Inspector) inspector.Inspector {
	wrappers := make([]inspector.Inspector, len(inspectors))

	for i, isp := range inspectors {
		wrappers[i] = func(node ast.Node) bool {
			file, isFile := node.(*ast.File)
			if !isFile {
				return false
			}

			for _, decl := range file.Decls {
				isp(decl)
			}

			return false
		}
	}

	return inspector.Lieutenant(wrappers...)
}
