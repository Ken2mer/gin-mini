package gin

import (
	"net/http"
	"sync"
)

// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
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
	// c.writermem.reset(w)
	c.Request = req
	// c.reset()

	// engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}
