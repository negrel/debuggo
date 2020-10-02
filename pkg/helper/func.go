package helper

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/negrel/debuggo/pkg/inspector"
)

// Func is an helper to edit functions.
type Func struct {
	filter func(name string) (replaceName string, ok bool)
}

// NewFunc return an helper to edit functions.
func NewFunc() *Func {
	return &Func{}
}

func (f *Func) RenameFunc(filter func(name string) (replaceName string, ok bool)) (renameFuncDecl, renameFuncCall inspector.Inspector) {
	f.filter = filter

	return f.renameFuncDecl, f.renameFuncCall
}

func (f *Func) renameFuncDecl(node ast.Node) (recursive bool) {
	recursive = true

	funcDecl, isFuncDecl := node.(*ast.FuncDecl)
	if !isFuncDecl {
		return
	}

	funcName := funcDecl.Name.Name

	newName, ok := f.filter(funcName)
	if !ok || newName == "" {
		return false
	}
	funcDecl.Name.Name = newName

	return
}

func (f *Func) renameFuncCall(node ast.Node) (recursive bool) {
	recursive = true

	callExpr, isCallExpr := node.(*ast.CallExpr)
	if !isCallExpr {
		return
	}

	switch fun := callExpr.Fun.(type) {
	case *ast.Ident:
		newName, ok := f.filter(fun.Name)
		if !ok {
			return
		}
		f.replaceFunInCallExpr(callExpr, newName)

	default:
		return
	}

	return
}

func (f *Func) replaceFunInCallExpr(callExpr *ast.CallExpr, newName string) {
	split := strings.Split(newName, ".")

	if length := len(split); length == 2 {
		callExpr.Fun = &ast.SelectorExpr{
			Sel: ast.NewIdent(split[0]),
			X:   ast.NewIdent(split[1]),
		}
	} else if length == 1 {
		callExpr.Fun = ast.NewIdent(newName)
	} else {
		panic(fmt.Sprintf("%v is an invalid new name", newName))
	}
}
