package testUtils

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/cache"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/router"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// Global test variables
var Handler router.Handler
var CacheConn redis.Conn

// Utility methods
func init() {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	engine.NoRoute(router.NoRouteHandler())
	engine.Use(router.AppRequestHandler())
	Handler = router.Handler{Engine: engine}
	Handler.SetupApiEndpoints()

	CacheConn = cache.InitForMockTest()
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
