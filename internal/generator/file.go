package generator

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"io/ioutil"
	"reflect"
	"strings"

	"fmt"
)

// File define a parsed file.
type file struct {
	pkg  *Package
	ast  *ast.File
	name string
	dir  string
}

func (f *file) generateContent(release bool) (content []byte, err error) {
	buf := bytes.Buffer{}
	buf.WriteString("// Code generated by Debuggo. DO NOT EDIT.\n\n")
	buf.WriteString("//go:generate debuggo " + f.pkg.path + "\n")

	if release {
		buf.WriteString("\n// +build !" + f.pkg.name + "\n\n")
		err = f.generateReleaseContent(&buf)
	} else {
		buf.WriteString("\n// +build " + f.pkg.name + "\n\n")
		err = f.generateDebugContent(&buf)
	}

	// Generate and format
	if err != nil {
		return buf.Bytes(), err
	}

	content, err = format.Source(buf.Bytes())
	if err != nil {
		return content, err
	}

	return
}

func (f *file) generateDebugContent(buf *bytes.Buffer) error {
	path := f.dir + "/" + f.name
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	buf.Write(src)

	return nil
}

func (f *file) generateReleaseContent(buf *bytes.Buffer) error {
	// Package declaration
	buf.WriteString("package " + f.pkg.name + "\n\n")

	// Getting signature for all top-level declaration
	signatures := make([]string, 0)
	var afile *ast.File

	// Getting the file object
	// that contain top level declaration.
	ast.Inspect(f.ast, func(n ast.Node) bool {

		// fmt.Print(reflect.TypeOf(n))
		// fmt.Printf(" %+v\n", n)
		// return true

		// File object contain a top-level declaration field.
		switch node := n.(type) {
		case *ast.File:
			afile = node

		// Looking for identifier that point to an imported package.
		case *ast.Ident:
			fmt.Printf("%+v\n", node)

		default:
			fmt.Print(reflect.TypeOf(node))
			fmt.Printf(" %+v\n", node)
		}

		return false
	})

	path := f.dir + "/" + f.name
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	for _, decl := range afile.Decls {
		if fDecl, ok := decl.(*ast.FuncDecl); ok {
			fType := declToString(fDecl.Type)
			signatures = append(signatures, fmt.Sprintf("%v%v", fDecl.Name, fType))
			continue
		}

		// We don't care about imports.
		gDecl, ok := decl.(*ast.GenDecl)
		if !ok || gDecl.Tok == token.IMPORT {
			continue
		}

		// copying 'const', 'var' and 'type' declaration
		start, end := f.pkg.fset.Position(gDecl.Pos()), f.pkg.fset.Position(gDecl.End())

		copy := copyBytes(src, start, end)
		copy = append(copy, '\n')
		buf.Write(copy)
	}

	// Writing empty function based on signature
	for _, fct := range signatures {
		buf.WriteString(fmt.Sprintf("func %v {}\n", fct))
	}

	return nil
}

func copyBytes(src []byte, from, to token.Position) []byte {
	line := 1
	start, end := -1, -1

	for i, bbyte := range src {
		if bbyte == '\n' {
			line++
		}

		if start == -1 && line == from.Line {
			start = i + from.Column
		}

		if end == -1 && line == to.Line {
			end = i + to.Column
		}
	}

	return src[start:end]
}

// return the signature of a function declaration
func declToString(e ast.Expr) string {
	switch expr := e.(type) {
	case *ast.Ellipsis:
		subType := declToString(expr.Elt)
		return fmt.Sprintf("...%s", subType)

	case *ast.InterfaceType:
		fNames := fmtFieldList("%v%v", expr.Methods.List)
		return fmt.Sprintf("interface{%v}", strings.Join(fNames, "\n"))

	case *ast.SelectorExpr:
		return fmt.Sprintf("%v.%v", expr.X, expr.Sel)

	case *ast.FuncType:
		p := []string{}
		if expr.Params != nil && expr.Params.List != nil {
			p = fmtFieldList("%v %v", expr.Params.List)
		}

		params := strings.Join(p, ", ")

		r := []string{}
		if expr.Results != nil && expr.Results.List != nil {
			r = fmtFieldList("%v %v", expr.Results.List)
		}

		results := ""
		if len := len(r); len != 0 {
			results = r[0]
			if len > 1 {
				results = fmt.Sprintf("(%v)", strings.Join(r, ", "))
			}
		}

		return fmt.Sprintf("(%v) %v", params, results)

	case *ast.ArrayType:
		arrType := declToString(expr.Elt)
		len := ""
		// Array size (len is nil for slices)
		if expr.Len != nil {
			// Const
			ident, ok := expr.Len.(*ast.Ident)
			if ok {
				decl := ident.Obj.Decl.(*ast.ValueSpec)
				lit := decl.Values[0].(*ast.BasicLit)
				len = fmt.Sprint(lit.Value)
			}

			// Literal
			lit, ok := expr.Len.(*ast.BasicLit)
			if ok {
				len = lit.Value
			}
		}

		return fmt.Sprintf("[%v]%v", len, arrType)

	// primitive
	default:
		return fmt.Sprint(expr)
	}

}

// Format the given field list with the given format string.
func fmtFieldList(format string, fieldsList []*ast.Field) []string {
	fields := make([]string, 0, len(fieldsList))
	for _, field := range fieldsList {
		names := []string{}
		for _, name := range field.Names {
			names = append(names, name.Name)
		}

		ftype := declToString(field.Type)

		// Avoid extra space for nameless fields (e.g returns)
		if len(names) != 0 {
			fields = append(fields, fmt.Sprintf(format, strings.Join(names, ", "), ftype))
		} else {
			fields = append(fields, fmt.Sprintf(format, ftype, ""))
		}
	}

	return fields
}
