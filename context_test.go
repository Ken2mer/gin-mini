package gin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextReset(t *testing.T) {
	router := New()
	c := router.allocateContext()
	assert.Equal(t, c.engine, router)

	c.index = 2
	c.Writer = &responseWriter{ResponseWriter: httptest.NewRecorder()}
	// c.Params = Params{Param{}}
	// c.Error(errors.New("test")) // nolint: errcheck
	// c.Set("foo", "bar")
	c.reset()

	// assert.False(t, c.IsAborted())
	// assert.Nil(t, c.Keys)
	// assert.Nil(t, c.Accepted)
	// assert.Len(t, c.Errors, 0)
	// assert.Empty(t, c.Errors.Errors())
	// assert.Empty(t, c.Errors.ByType(ErrorTypeAny))
	// assert.Len(t, c.Params, 0)
	assert.EqualValues(t, c.index, -1)
	assert.Equal(t, c.Writer.(*responseWriter), &c.writermem)
}

func TestContextRenderJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	// c.JSON(http.StatusCreated, H{"foo": "bar", "html": "<b>"})
	c.JSON(http.StatusOK, H{"foo": "bar", "html": "<b>"})

	// assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"foo\":\"bar\",\"html\":\"\\u003cb\\u003e\"}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}
