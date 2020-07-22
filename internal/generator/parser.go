package generator

import (
	"go/ast"
	"go/token"
	"strings"
)

type parser struct {
	requiredPkg  map[string]string
	potentialPkg map[string]string
}

func newParser() *parser {
	return &parser{
		requiredPkg:  make(map[string]string, 8),
		potentialPkg: make(map[string]string, 8),
	}
}

// Parse the given ast.file and edit it to remove useless
// import and top-level function body.
func (p *parser) parse(f *ast.File) {
	for _, d := range f.Decls {
		switch decl := d.(type) {
		case *ast.FuncDecl:
			// We remove the body of the function
			decl.Body.List = []ast.Stmt{}
			// Parse to find some required package in function type.
			p.parseExpr(decl.Type)

		case *ast.GenDecl:
			if decl.Tok == token.IMPORT {
				// Storing imports in map for checking
				for _, spec := range decl.Specs {
					importSpec := spec.(*ast.ImportSpec)

					var name, path string
					// import path
					if importSpec.Path != nil {
						path = importSpec.Path.Value
					}

					// Name if custome package name
					if importSpec.Name != nil {
						name = importSpec.Name.Name
					} else {
						s := strings.Split(path, "/")
						name = strings.Trim(s[len(s)-1], "\"")
					}

					p.potentialPkg[name] = path
				}
				// Delete unused imports from the AST.
				defer func() {
					specs := make([]ast.Spec, 0, len(p.potentialPkg))

					for _, spec := range decl.Specs {
						importSpec := spec.(*ast.ImportSpec)
						if importSpec.Path == nil {
							continue
						}

						var key string
						if importSpec.Name != nil {
							key = importSpec.Name.Name
						} else {
							s := strings.Split(importSpec.Path.Value, "/")
							key = strings.Trim(s[len(s)-1], "\"")
						}

						if p.requiredPkg[key] == importSpec.Path.Value {
							specs = append(specs, importSpec)
						}
					}

					decl.Specs = specs
				}()
				continue
			}

			if decl.Tok == token.CONST || decl.Tok == token.VAR {
				for _, spec := range decl.Specs {
					valueSpec := spec.(*ast.ValueSpec)

					// Parse to find some required package in var/const declaration.
					for _, value := range valueSpec.Values {
						p.parseExpr(value)
					}
				}
			}
		}
	}
}

func (p *parser) parseExpr(expr ast.Expr) {
	// Checking for package required in function signatures.
	switch e := expr.(type) {
	// We only care about SelectorExpr but some expression
	// can contain selector so we parse them recursively.
	case *ast.SelectorExpr:
		if e.X != nil {
			x, ok := e.X.(*ast.Ident)
			if !ok {
				return
			}

			if path := p.potentialPkg[x.String()]; path != "" {
				p.requiredPkg[x.String()] = path
			}
		}

	case *ast.FuncType:
		// PARAMS
		if e.Params != nil {
			for _, param := range e.Params.List {
				// Rename all parameters to "_"
				for _, name := range param.Names {
					name.Name = "_"
				}

				// Parameter Type can be a SelectorExpr (e.g io.Writer)
				p.parseExpr(param.Type)
			}
		}

		// RESULTS
		if e.Results != nil {
			for _, result := range e.Results.List {
				// Rename all results to "_"
				for _, name := range result.Names {
					name.Name = "_"
				}

				// Result Type can be a SelectorExpr (e.g io.Writer)
				p.parseExpr(result.Type)
			}
		}

	case *ast.ArrayType:
		p.parseExpr(e.Elt)

	case *ast.InterfaceType:
		// Interface can contain an interface from another
		// package (e.g fmt.Stringer)
		for _, method := range e.Methods.List {
			p.parseExpr(method.Type)
		}

		// default:
		// 	fmt.Println(reflect.TypeOf(e))
		// 	fmt.Printf("%+v\n", e)
	}
}
