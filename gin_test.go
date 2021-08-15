package gin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestCreateEngine(t *testing.T) {
	router := New()
	// assert.Equal(t, "/", router.basePath)
	assert.Equal(t, router.engine, router)
	assert.Empty(t, router.Handlers)
}

func TestAddRoute(t *testing.T) {
	router := New()
	router.addRoute("GET", "/", HandlersChain{func(_ *Context) {}})

	assert.Len(t, router.trees, 1)
	assert.NotNil(t, router.trees.get("GET"))
	assert.Nil(t, router.trees.get("POST"))

	router.addRoute("POST", "/", HandlersChain{func(_ *Context) {}})

	assert.Len(t, router.trees, 2)
	assert.NotNil(t, router.trees.get("GET"))
	assert.NotNil(t, router.trees.get("POST"))

	router.addRoute("POST", "/post", HandlersChain{func(_ *Context) {}})
	assert.Len(t, router.trees, 2)
}

func TestEngineHandleContext(t *testing.T) {
	r := New()
	r.GET("/", func(c *Context) {
		// c.Request.URL.Path = "/v2"
		r.HandleContext(c)
	})
	// v2 := r.Group("/v2")
	// {
	// 	v2.GET("/", func(c *Context) {})
	// }

	assert.NotPanics(t, func() {
		// w := performRequest(r, "GET", "/")
		_ = performRequest(r, "GET", "/")
		// assert.Equal(t, 301, w.Code)
	})
}
