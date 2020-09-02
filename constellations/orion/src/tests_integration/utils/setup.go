package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/middlewares"
)

var ENV_VAR_TEST = "TEST_ENV"
var ENV_VAR_CIRCLE_CI = "test_ci"

func SetupTestEnvironment(m *testing.M) {
	fmt.Println("Setting up Test Environment...")

	fmt.Println("Setting up test logger")
	logger.SetupTest()

	logger.Message("Retrieving configurations...")
	var configPath string
	if os.Getenv(ENV_VAR_TEST) == ENV_VAR_CIRCLE_CI {
		configPath = "./configs/ci.yml"
	} else {
		configPath = "./configs/local.yml"
	}
	config := middlewares.RetrieveConfigurations(configPath)

	logger.Message("Connecting to database...")
	configDb := config.Database
	SetupTestDatabase(configDb.Host, configDb.Port, configDb.Username, configDb.Password, configDb.DbName)
	defer CloseDb()

	logger.Message("Setting up caching...")
	configCache := config.Cache
	SetupTestCache(configCache.Host, configCache.Port, configCache.Password)
	defer CloseCache()

	logger.Message("Setting up router...")
	SetupTestRouter()
	os.Exit(m.Run())
}
