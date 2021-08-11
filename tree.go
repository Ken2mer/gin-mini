package gin

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	// Key   string
	// Value string
}

// Params is a Param-slice, as returned by the router.
type Params []Param

type methodTree struct {
	// method string
	// root   *node
}

type methodTrees []methodTree
