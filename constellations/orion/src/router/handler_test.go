package router_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/cache"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Global test variables
var Handler router.Handler

// Initializer
func init() {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	engine.NoRoute(router.NoRouteHandler())
	engine.Use(router.AppRequestHandler())
	Handler = router.Handler{Engine: engine}
	Handler.SetupApiEndpoints()

	cache.InitForMockTest()
}

// Utility method
func sendHttpRequest(t *testing.T, method, url string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Errorf("http request error: %v\n", err)
	}
	w := httptest.NewRecorder()
	Handler.Engine.ServeHTTP(w, req)
	return w
}

func TestGetHealth(t *testing.T) {
	recorder := sendHttpRequest(t, http.MethodGet, "/api/health", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestPathNotFound(t *testing.T) {
	recorder := sendHttpRequest(t, http.MethodGet, "/api/endpoint-doesnt-exist", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}