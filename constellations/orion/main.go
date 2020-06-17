package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/middlewares"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/router"
)

func main() {
	fmt.Println("Orion service starting...")

	// App Repos
	fmt.Println("Setting up Repos...")
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDefault := os.Getenv("DB_DEFAULT")

	db := repos.Open(dbHost, dbPort, dbUser, dbPassword, dbDefault)
	repos.Migrate(db, "file://src/repos/migrations")
	repos.SetupRepos(db)
	defer repos.Close(db)
	fmt.Println("Database started!")

	// App Router
	fmt.Println("Setting up Router...")
	engine := gin.Default()
	fmt.Println("Setting up Middlewares...")
	corsOriginEnvVar := os.Getenv("CORS_ORIGIN")
	corsOrigins := []string{corsOriginEnvVar}
	configCors := middlewares.CreateCorsConfig(corsOrigins)
	engine.Use(cors.New(configCors))
	handler := router.Handler{Engine: engine}
	handler.SetupApiEndpoints()

	// Run web server
	handler.Engine.Run(":8001")
}
