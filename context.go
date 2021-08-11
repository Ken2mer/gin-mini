package gin

import "net/http"

// Context is the most important part of gin.
type Context struct {
	// writermem responseWriter
	Request *http.Request
	// Writer    ResponseWriter

	// Params   Params
	// handlers HandlersChain
	// index    int8
	// fullPath string

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
