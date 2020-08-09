package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/router"
	"github.com/gin-gonic/gin"
)

var handler *router.Handler

func SetupTestRouter() {
	logger.Message("Initializing Router...")
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	newHandler := router.Handler{Engine: engine}
	newHandler.SetupApiEndpoints()
	handler = &newHandler
}

func SendHttpRequest(t *testing.T, method, url string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Errorf("Error: http request error: %v\n", err)
	}

	if handler == nil {
		t.Fatalf("Error: http handler is nil")
	}

	w := httptest.NewRecorder()
	handler.Engine.ServeHTTP(w, req)
	return w
}

func CreateJsonBody(v interface{}) io.Reader {
	marshal, _ := json.Marshal(v)
	return bytes.NewBuffer(marshal)
}
