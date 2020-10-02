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

// New return an inspector Lead.
func New(inspectors ...Inspector) *Lead {
	return &Lead{
		inspectors: inspectors,
		active:     len(inspectors),
	}
}

// Inspect start the inspection of the given node.
func (l *Lead) Inspect(node ast.Node) {
	ast.Inspect(node, l.inspect)
}

func (l *Lead) inspect(node ast.Node) bool {
	for index, inspector := range l.inspectors {
		if l.active == 0 {
			l.enableAllInspector()
			return false
		}

		notRecursiveHook := !inspector(node)

		if notRecursiveHook {
			l.disableInspectorUntilNextTree(index, inspector)
		}
	}

	return true
}

func (l *Lead) disableInspectorUntilNextTree(index int, inspector Inspector) {
	l.active--

	l.inspectors[index] = func(n ast.Node) bool {
		if n == nil {
			l.inspectors[index] = inspector
			l.active++
		}

		return true
	}
}

func (l *Lead) enableAllInspector() {
	for _, inspector := range l.inspectors {
		inspector(nil)
	}
}
