package gin

import (
	"net/http"
	"sync"
)

// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

type Engine struct {
	RouterGroup

	pool      sync.Pool
	trees     methodTrees
	maxParams uint16
}

// New returns a new blank Engine instance without any middleware attached.
func New() *Engine {
	// debugPrintWARNINGNew()
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			// basePath: "/",
			// root: true,
		},

		trees: make(methodTrees, 0, 9),
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
	// debugPrintWARNINGDefault()
	engine := New()
	// engine.Use(Logger())
	// engine.Use(Recovery())
	return engine
}

func (engine *Engine) allocateContext() *Context {
	v := make(Params, 0, engine.maxParams)
	return &Context{engine: engine, params: &v}
}

func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
	// assert1(path[0] == '/', "path must begin with '/'")
	// assert1(method != "", "HTTP method can not be empty")
	// assert1(len(handlers) > 0, "there must be at least one handler")

	// debugPrintRoute(method, path, handlers)

	root := engine.trees.get(method)
	if root == nil {
		root = new(node)
		// root.fullPath = "/"
		engine.trees = append(engine.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handlers)

	// Update maxParams
	// if paramsCount := countParams(path); paramsCount > engine.maxParams {
	// 	engine.maxParams = paramsCount
	// }
}

// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
func (engine *Engine) Run(addr ...string) (err error) {
	// defer func() { debugPrintError(err) }()

	// err = engine.parseTrustedProxies()
	// if err != nil {
	// 	return err
	// }

	address := resolveAddress(addr)
	// debugPrint("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine)
	return
}

// ServeHTTP conforms to the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)
	c.writermem.reset(w)
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}

func (engine *Engine) handleHTTPRequest(c *Context) {
	// httpMethod := c.Request.Method
	rPath := c.Request.URL.Path
	unescape := false

	// Find root of the tree for the given HTTP method
	t := engine.trees
	for i, tl := 0, len(t); i < tl; i++ {
		// if t[i].method != httpMethod {
		// 	continue
		// }
		root := t[i].root
		// Find route in tree
		value := root.getValue(rPath, c.params, unescape)
		// if value.params != nil {
		// 	c.Params = *value.params
		// }
		if value.handlers != nil {
			c.handlers = value.handlers
			// c.fullPath = value.fullPath
			c.Next()
			// c.writermem.WriteHeaderNow()
			return
		}
		break
	}

	// c.handlers = engine.allNoRoute
	// serveError(c, http.StatusNotFound, default404Body)
}
