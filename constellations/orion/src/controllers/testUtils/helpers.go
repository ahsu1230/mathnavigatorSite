package testUtils

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/router"
	"github.com/gin-gonic/gin"
)

// Global test variables
var Handler router.Handler

// Utility methods
func init() {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	Handler = router.Handler{Engine: engine}
	Handler.SetupApiEndpoints()
}

func SendHttpRequest(t *testing.T, method, url string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Errorf("http request error: %v\n", err)
	}
	w := httptest.NewRecorder()
	Handler.Engine.ServeHTTP(w, req)
	return w
}