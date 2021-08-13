package gin

import (
	"net/http"
)

// IRouter defines all router handle interface includes single and group router.
type IRouter interface {
	IRoutes
	// Group(string, ...HandlerFunc) *RouterGroup
}

// IRoutes defines all router handle interface.
type IRoutes interface {
	Use(...HandlerFunc) IRoutes
}

// RouterGroup is used internally to configure router
type RouterGroup struct {
	Handlers HandlersChain
	// basePath string
	engine *Engine
	// root   bool
}

var _ IRouter = &RouterGroup{}

// Use adds middleware to the group, see example code in GitHub.
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	group.Handlers = append(group.Handlers, middleware...)
	return group.returnObj()
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	// absolutePath := group.calculateAbsolutePath(relativePath)
	// handlers = group.combineHandlers(handlers)
	// group.engine.addRoute(httpMethod, absolutePath, handlers)
	group.engine.addRoute(httpMethod, relativePath, handlers)
	return group.returnObj()
}

// GET is a shortcut for router.Handle("GET", path, handle).
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodGet, relativePath, handlers)
}

func (group *RouterGroup) returnObj() IRoutes {
	// if group.root {
	// 	return group.engine
	// }
	return group
}
