package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/middlewares"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/repos"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/router"
)

func main() {
	fmt.Println("Orion service starting...")

	// App Configurations
	configFile := os.Args[1]
	config := middlewares.RetrieveConfigurations(configFile)
	fmt.Println("Building server in mode: ", config.App.Build)

	// App Repos
	fmt.Println("Setting up Repos...")
	configDb := config.Database
	db := repos.Open(configDb.Host, configDb.Port, configDb.Username, configDb.Password, configDb.DbName)
	repos.Migrate(db, "file://pkg/repos/migrations")
	repos.SetupRepos(db)
	defer repos.Close(db)
	fmt.Println("Database started!")

	// App Router
	fmt.Println("Setting up Router...")
	engine := gin.Default()
	fmt.Println("Setting up Middlewares...")
	corsOrigins := []string{config.App.CorsOrigin}
	configCors := middlewares.CreateCorsConfig(corsOrigins)
	engine.Use(cors.New(configCors))
	handler := router.Handler{Engine: engine}
	handler.SetupApiEndpoints()

	// Run web server
	handler.Engine.Run(":8080")
}
