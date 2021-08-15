package gin

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testRouteOK(method string, t *testing.T) {
	passed := false
	// passedAny := false
	r := New()
	// r.Any("/test2", func(c *Context) {
	// 	passedAny = true
	// })
	// r.Handle(method, "/test", func(c *Context) {
	// 	passed = true
	// })
	r.GET("/test", func(c *Context) {
		passed = true
	})

	w := performRequest(r, method, "/test")
	assert.True(t, passed)
	assert.Equal(t, http.StatusOK, w.Code)

	// performRequest(r, method, "/test2")
	// assert.True(t, passedAny)
}

func TestRouterGroupRouteOK(t *testing.T) {
	testRouteOK(http.MethodGet, t)
	// testRouteOK(http.MethodPost, t)
	// testRouteOK(http.MethodPut, t)
	// testRouteOK(http.MethodPatch, t)
	// testRouteOK(http.MethodHead, t)
	// testRouteOK(http.MethodOptions, t)
	// testRouteOK(http.MethodDelete, t)
	// testRouteOK(http.MethodConnect, t)
	// testRouteOK(http.MethodTrace, t)
}
