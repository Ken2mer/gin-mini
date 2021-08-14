package gin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
