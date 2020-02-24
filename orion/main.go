package main
import (
    "fmt"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"

    "github.com/ahsu1230/mathnavigatorSite/orion/middlewares"
    "github.com/ahsu1230/mathnavigatorSite/orion/repos"
    "github.com/ahsu1230/mathnavigatorSite/orion/router"
)

func main() {
    fmt.Println("Orion service starting...")

    // App Configurations
    config := middlewares.RetrieveConfigurations()
    fmt.Println("Building server in mode: ", config.App.Build)

    // App Repos
    fmt.Println("Setting up Repos...")
    configDb := config.Database
    db := repos.Open(configDb.Host, configDb.Port, configDb.Username, configDb.Password)
    repos.Migrate(db)
    repos.SetupRepos(db)
    defer repos.Close(db)
    fmt.Println("Database started!")

    // App Router
    fmt.Println("Setting up Router...")
    engine := gin.Default()
    fmt.Println("Setting up Middlewares...")
    configCors := middlewares.CreateCorsConfig(config);
    engine.Use(cors.New(configCors))
    handler := router.Handler { engine }
    handler.SetupApiEndpoints()

    // Run web server
    handler.Engine.Run(":8080")
}
