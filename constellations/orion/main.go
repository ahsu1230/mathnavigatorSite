package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/middlewares"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	repoUtils "github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/router"
)

func main() {
	fmt.Println("Orion service starting...")

	// Setup Logging
	fmt.Println("Setting up Logger...")
	logger.SetupDev()
	if production {
		err := logger.SetupProd()
		if err != nil {
			fmt.Fatalf("Logger failed to setup! %w", err)
			return
		}
	}
	logger.Message("Logger successfully setup!")

	// App Repos
	logger.Message("Setting up Repos...")
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDefault := os.Getenv("DB_DEFAULT")

	db := repoUtils.Open(dbHost, dbPort, dbUser, dbPassword, dbDefault)
	repoUtils.Migrate(db, "file://src/repos/migrations")
	repos.SetupRepos(db)
	defer repoUtils.Close(db)
	logger.Message("Database started!")

	// App Router
	logger.Message("Setting up Router...")
	engine := gin.Default()
	
	logger.Message("Setting up Middlewares...")
	corsOriginEnvVar := os.Getenv("CORS_ORIGIN")
	corsOrigins := []string{corsOriginEnvVar}
	configCors := middlewares.CreateCorsConfig(corsOrigins)
	engine.Use(cors.New(configCors))
	handler := router.Handler{Engine: engine}
	handler.SetupApiEndpoints()

	// Run web server
	handler.Engine.Run(":8001")
}
