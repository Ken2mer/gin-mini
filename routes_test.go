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

func TestRouteParamsByName(t *testing.T) {
	name := ""
	lastName := ""
	wild := ""
	router := New()
	router.GET("/test/:name/:last_name/*wild", func(c *Context) {
		name = c.Params.ByName("name")
		lastName = c.Params.ByName("last_name")
		var ok bool
		wild, ok = c.Params.Get("wild")

		assert.True(t, ok)
		assert.Equal(t, name, c.Param("name"))
		assert.Equal(t, lastName, c.Param("last_name"))

		assert.Empty(t, c.Param("wtf"))
		assert.Empty(t, c.Params.ByName("wtf"))

		wtf, ok := c.Params.Get("wtf")
		assert.Empty(t, wtf)
		assert.False(t, ok)
	})

	w := performRequest(router, http.MethodGet, "/test/john/smith/is/super/great")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "john", name)
	assert.Equal(t, "smith", lastName)
	assert.Equal(t, "/is/super/great", wild)
}

func TestRouteContextHoldsFullPath(t *testing.T) {
	router := New()

	// Test routes
	routes := []string{
		"/simple",
		"/project/:name",
		"/",
		"/news/home",
		"/news",
		"/simple-two/one",
		"/simple-two/one-two",
		"/project/:name/build/*params",
		"/project/:name/bui",
		"/user/:id/status",
		"/user/:id",
		"/user/:id/profile",
	}

	for _, route := range routes {
		actualRoute := route
		router.GET(route, func(c *Context) {
			// For each defined route context should contain its full path
			assert.Equal(t, actualRoute, c.FullPath())
			// c.AbortWithStatus(http.StatusOK)
		})
	}

	for _, route := range routes {
		w := performRequest(router, http.MethodGet, route)
		assert.Equal(t, http.StatusOK, w.Code)
	}

	// Test not found
	// router.Use(func(c *Context) {
	// 	// For not found routes full path is empty
	// 	assert.Equal(t, "", c.FullPath())
	// })

	// w := performRequest(router, http.MethodGet, "/not-found")
	// assert.Equal(t, http.StatusNotFound, w.Code)
}
