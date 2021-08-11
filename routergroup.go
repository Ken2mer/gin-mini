package gin

// RouterGroup is used internally to configure router
type RouterGroup struct {
	Handlers HandlersChain
	// basePath string
	engine *Engine
	// root   bool
}
