package gin

import (
	"net/http"

	"github.com/Ken2mer/gin/render"
)

// Context is the most important part of gin.
type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	// Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine *Engine
	params *Params

	// mu sync.RWMutex

	// Keys is a key/value pair exclusively for the context of each request.
	// Keys map[string]interface{}

	// Errors is a list of errors attached to all the handlers/middlewares who used this context.
	// Errors errorMsgs

	// Accepted defines a list of manually accepted formats for content negotiation.
	// Accepted []string

	// queryCache url.Values

	// formCache url.Values

	// sameSite http.SameSite
}

func (c *Context) reset() {
	c.Writer = &c.writermem
	// c.Params = c.Params[:0]
	c.handlers = nil
	c.index = -1

	// c.fullPath = ""
	// c.Keys = nil
	// c.Errors = c.Errors[:0]
	// c.Accepted = nil
	// c.queryCache = nil
	// c.formCache = nil
	// *c.params = (*c.params)[:0]
}

// FullPath returns a matched route full path.
func (c *Context) FullPath() string {
	return c.fullPath
}

// Next executes the pending handlers in the chain inside the calling handler.
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// Render writes the response headers and calls render.Render to render data.
func (c *Context) Render(code int, r render.Render) {
	// c.Status(code)

	// if !bodyAllowedForStatus(code) {
	// 	r.WriteContentType(c.Writer)
	// 	c.Writer.WriteHeaderNow()
	// 	return
	// }

	if err := r.Render(c.Writer); err != nil {
		panic(err)
	}
}

// JSON serializes the given struct as JSON into the response body.
func (c *Context) JSON(code int, obj interface{}) {
	c.Render(code, render.JSON{Data: obj})
}
