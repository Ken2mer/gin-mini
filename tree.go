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

type nodeType uint8

// const (
// 	static nodeType = iota // default
// 	root
// 	param
// 	catchAll
// )

type node struct {
	path string
	// indices   string
	wildChild bool
	nType     nodeType
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

type nodeValue struct {
	handlers HandlersChain
	// params   *Params
	// tsr bool
	// fullPath string
}

func (n *node) getValue(path string, params *Params, unescape bool) (value nodeValue) {

	// walk: // Outer loop for walking the tree
	for {
		prefix := n.path
		if len(path) > len(prefix) {
			if path[:len(prefix)] == prefix {
				path = path[len(prefix):]

				// If there is no wildcard pattern, recommend a redirection
				if !n.wildChild {
					// Nothing found.
					// We can recommend to redirect to the same URL without a
					// trailing slash if a leaf exists for that path.
					// value.tsr = path == "/" && n.handlers != nil
					return
				}

				// Handle wildcard child, which is always at the end of the array
				// n = n.children[len(n.children)-1]

				switch n.nType {
				default:
					panic("invalid node type")
				}
			}
		}

		if path == prefix {
			if value.handlers = n.handlers; value.handlers != nil {
				// value.fullPath = n.fullPath
				return
			}

			return
		}

		return
	}
}
