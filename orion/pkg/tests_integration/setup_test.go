package integration_tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/middlewares"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/router"
)

var db *sql.DB
var handler router.Handler

func TestMain(m *testing.M) {
	fmt.Println("Setting up Test Environment...")
	config := middlewares.RetrieveConfigurations("./config_integrations.yaml")
	configDb := config.Database
	db = repos.Open(configDb.Host, configDb.Port, configDb.Username, configDb.Password, "")	
	defer db.Close()
	setupTestDatabase(configDb.DbName)
	setupTestRouter()
	os.Exit(m.Run())
}

// Helper methods for Database
func setupTestDatabase(dbName string) {

	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}
	tx.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName))
	tx.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	tx.Exec(fmt.Sprintf("USE %s;", dbName))
	tx.Commit()

	// // Drop database if already exists
	// fmt.Println("Dropping if exists ", dbName)
	// if _, err := db.Exec(); err != nil {
	// 	panic(err.Error())
	// }

	// // Create database
	// fmt.Println("Creating database ", dbName)
	// if _, err := db.Exec(); err != nil {
	// 	panic(err.Error())
	// }
	
	// // Use database
	// fmt.Println("Using database ", dbName)
	// if _, err := db.Exec(); err != nil {
	// 	panic(err.Error())
	// }
	fmt.Println("Starting migrations...")
	repos.Migrate(db, "file://../repos/migrations")
	
	fmt.Println("Initializing repos...")
	repos.ProgramRepo.Initialize(db)
	// Initialize other tables here...

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}
}

func refreshTable(t *testing.T, tableName string) error {
	stmt, err := db.Prepare(fmt.Sprintf("TRUNCATE TABLE %s;", tableName))
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		t.Fatalf("Error truncating tables: %s", err)
	}
	return nil
}

// Helper methods for Router
func setupTestRouter() {
	fmt.Println("Initializing Router...")
	gin.SetMode(gin.TestMode)
    engine := gin.Default()
	handler = router.Handler{ Engine: engine }
    handler.SetupApiEndpoints()
}

func sendHttpRequest(t *testing.T, method, url string, body io.Reader) *httptest.ResponseRecorder {
    req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Errorf("http request error: %v\n", err)
    }
    w := httptest.NewRecorder()
    handler.Engine.ServeHTTP(w, req)
    return w
}

func createJsonBody(v interface{}) io.Reader {
	marshal, _ := json.Marshal(v)
	return bytes.NewBuffer(marshal)
}