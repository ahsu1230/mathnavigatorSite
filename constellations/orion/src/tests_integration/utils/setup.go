package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/middlewares"
)

func SetupTestEnvironment(m *testing.M) {
	fmt.Println("Setting up Test Environment...")

	var configPath string
	if os.Getenv("TEST_ENV") == "test_ci" {
		configPath = "./configs/ci.yml"
	} else {
		configPath = "./configs/local.yml"
	}
	config := middlewares.RetrieveConfigurations(configPath)

	fmt.Println("Connecting to database...")
	configDb := config.Database
	SetupTestDatabase(configDb.Host, configDb.Port, configDb.Username, configDb.Password, configDb.DbName)
	defer db.Close()

	fmt.Println("Setting up router...")
	SetupTestRouter()
	os.Exit(m.Run())
}
