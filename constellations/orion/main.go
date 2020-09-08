package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/middlewares"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/cache"
	repoUtils "github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
	_ "github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/migrations1"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/router"
)

func main() {
	fmt.Println("Orion service starting...")

	// Parse flags
	var production bool
	flag.BoolVar(&production, "production", false, "Production mode")
	flag.Parse()

	// Setup Logging
	fmt.Println("Setting up Logger...")
	if production {
		err := logger.SetupProd()
		if err != nil {
			fmt.Printf("Logger failed to setup! %w", err)
			os.Exit(1)
			return
		}
	} else {
		logger.SetupDev()
	}
	logger.Message("Logger successfully setup!")

	// App Repos
	logger.Message("Setting up Repos...")
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDefault := os.Getenv("DB_DEFAULT")
	context := context.Background()

	db := repoUtils.Open(dbHost, dbPort, dbUser, dbPassword, dbDefault)
	repoUtils.Migrate(db, "file://src/repos/migrations")
	// repoUtils.Migrate1(db, "src/repos/migrations1")


	logger.Message("Invoking Migration command....")
	cmd := exec.Command("goose", "mysql", `"user:password@(db-mysql:3306)/mathnavdb?parseTime=true"`, "up")
	logger.Message(fmt.Sprintf("Command %s", cmd.String()))
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	logger.Message(fmt.Sprintf("Command finished. Exit Status %s", string(out)))
	

	repos.SetupRepos(context, db)
	defer repoUtils.Close(db)
	logger.Message("Database started!")

	logger.Message("Setting up Server-side caching...")
	cacheHost := os.Getenv("REDIS_HOST")
	cachePort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	cachePassword := os.Getenv("REDIS_PASSWORD")
	cache.Init(cacheHost, cachePort, cachePassword)
	defer cache.Close()

	// App Router
	logger.Message("Setting up Router...")
	engine := gin.Default()

	logger.Message("Setting up Middlewares...")
	corsOriginEnvVar := os.Getenv("CORS_ORIGIN")
	corsOrigins := []string{corsOriginEnvVar}
	configCors := middlewares.CreateCorsConfig(corsOrigins)
	engine.Use(cors.New(configCors))
	engine.NoRoute(router.NoRouteHandler())
	engine.Use(router.AppRequestHandler())
	handler := router.Handler{Engine: engine}
	handler.SetupApiEndpoints()

	// Run web server
	handler.Engine.Run(":8001")
}
