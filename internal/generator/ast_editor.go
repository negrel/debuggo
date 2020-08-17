package generator

import (
	"go/ast"
)

// FileEditor define any object that can edit an ast.File.
type FileEditor interface {
	edit(file *ast.File)
}

type nodeHook = func(n ast.Node) bool

type fileHook = func(f *ast.File)

type astEditor struct {
	nodeHooks       []nodeHook
	beforeEditHooks []fileHook
	afterEditHooks  []fileHook
}

func newAstEditor(options ...astEditorOption) FileEditor {
	e := &astEditor{
		nodeHooks:       make([]nodeHook, 0, 8),
		beforeEditHooks: make([]fileHook, 0, 8),
		afterEditHooks:  make([]fileHook, 0, 8),
	}

	for _, option := range options {
		option(e)
	}

	return e
}

func (e *astEditor) edit(file *ast.File) {
	for _, beforeEdit := range e.beforeEditHooks {
		beforeEdit(file)
	}

	ast.Inspect(file, func(n ast.Node) (recursive bool) {
		recursive = true

		// Simulate ast.Inspect call for every hook
		for hookIndex, nodeHook := range e.nodeHooks {
			notRecursiveHook := !nodeHook(n)

			if notRecursiveHook {
				e.disableNodeHookForNextSubTree(hookIndex, nodeHook)
			}
		}

		return
	})

	for _, afterEdit := range e.afterEditHooks {
		afterEdit(file)
	}
}

func (e *astEditor) disableNodeHookForNextSubTree(index int, hook nodeHook) {
	e.nodeHooks[index] = func(n ast.Node) bool {
		if n != nil {
			return true
		}

		e.nodeHooks[index] = hook

		return true
	}
}
