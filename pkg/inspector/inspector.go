package inspector

import (
	"go/ast"
)

type Inspector func(ast.Node) bool

// Chief is the inspector chief that manage the inspection.
type Chief struct {
	inspectors []Inspector
	active     int
}

// New return an inspector Chief.
func New(inspectors ...Inspector) *Chief {
	return &Chief{
		inspectors: inspectors,
		active:     len(inspectors),
	}
}

// Inspect start the inspection of the given node.
func (e *Chief) Inspect(node ast.Node) {
	ast.Inspect(node, func(n ast.Node) bool {
		for index, inspector := range e.inspectors {
			if e.active == 0 {
				e.enableAllInspector()
				return false
			}

			notRecursiveHook := !inspector(n)

			if notRecursiveHook {
				e.disableInspectorUntilNextTree(index, inspector)
			}
		}

		return true
	})
}

func (e *Chief) disableInspectorUntilNextTree(index int, inspector Inspector) {
	e.active--

	e.inspectors[index] = func(n ast.Node) bool {
		if n == nil {
			e.inspectors[index] = inspector
			e.active++
		}

		return true
	}
}

func (e *Chief) enableAllInspector() {
	for _, inspector := range e.inspectors {
		inspector(nil)
	}
}
