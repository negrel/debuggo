package debuggo

// Options define the debuggo generator options.
type Options struct {
	// Parse package option
	PkgPattern []string
	PkgTags    []string
	// Output is the relative path to output
	// debugging packages.
	Output string
}
