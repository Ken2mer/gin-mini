package gin

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	// Key   string
	// Value string
}

// Params is a Param-slice, as returned by the router.
type Params []Param

type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree

func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}

type node struct {
	path string
	// indices   string
	// wildChild bool
	// nType nodeType
	// priority uint32
	children []*node // child nodes, at most 1 :param style node at the end of the array
	handlers HandlersChain
	// fullPath string
}

func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	// n.priority++

	// Empty tree
	if len(n.path) == 0 && len(n.children) == 0 {
		n.insertChild(path, fullPath, handlers)
		// n.nType = root
		return
	}

// walk:
	for {
		return
	}
}

func (n *node) insertChild(path string, fullPath string, handlers HandlersChain) {

	// If no wildcard was found, simply insert the path and handle
	n.path = path
	n.handlers = handlers
	// n.fullPath = fullPath
}
