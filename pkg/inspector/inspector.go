package inspector

import (
	"go/ast"
)

type Inspector func(ast.Node) bool

// Lieutenant define an Inspector that manage his own Inspectors.
func Lieutenant(Inspectors ...Inspector) Inspector {
	return New(Inspectors...).inspect
}

// Lead is the Inspector chief that manage the inspection.
type Lead struct {
	inspectors map[int]Inspector
	depth      int
	inactive   map[int]map[int]Inspector
}

// New return an Inspector Lead.
func New(inspectors ...Inspector) *Lead {
	insp := make(map[int]Inspector, len(inspectors))
	for index, inspector := range inspectors {
		insp[index] = inspector
	}

	return &Lead{
		inspectors: insp,
		depth:      0,
		inactive:   make(map[int]map[int]Inspector),
	}
}

// Inspect start the inspection of the given node.
func (l *Lead) Inspect(node ast.Node) {
	ast.Inspect(node, l.inspect)
}

func (l *Lead) inspect(node ast.Node) bool {
	if node != nil {
		defer func() { l.depth++ }()
	} else {
		defer func() {
			l.depth--
			l.enableInactive()
		}()
	}

	for index, inspector := range l.inspectors {
		recursiveHook := inspector(node)

		if !recursiveHook {
			l.disableForSubTree(index)
		}
	}

	if len(l.inspectors) == 0 {
		l.enableInactive()
		return false
	}
	return true
}

func (l *Lead) disableForSubTree(index int) {
	if l.inactive[l.depth] == nil {
		l.inactive[l.depth] = make(map[int]Inspector)
	}

	l.inactive[l.depth][index] = l.inspectors[index]
	delete(l.inspectors, index)
}

func (l *Lead) enableInactive() {
	if _, ok := l.inactive[l.depth]; !ok {
		return
	}

	for index, inspector := range l.inactive[l.depth] {
		l.inspectors[index] = inspector
	}
	delete(l.inactive, l.depth)
}
