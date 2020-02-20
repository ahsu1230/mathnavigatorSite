package main
import (
    "fmt"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"

    "github.com/ahsu1230/mathnavigatorSite/orion/controllers"
    "github.com/ahsu1230/mathnavigatorSite/orion/middlewares"
    "github.com/ahsu1230/mathnavigatorSite/orion/router"
    "github.com/ahsu1230/mathnavigatorSite/orion/stores"
)

func main() {
    fmt.Println("Orion service starting...")

    // App Configurations
    config := middlewares.RetrieveConfigurations()
    fmt.Println("Building server in mode: ", config.App.Build)

    // App Store
    fmt.Println("Setting up Store...")
    configDb := config.Database
    dbSql, dbSqlx := stores.Open(configDb.Host, configDb.Port, configDb.Username, configDb.Password)
    stores.Migrate(dbSql)
    defer stores.Close(dbSql, dbSqlx)

    // App Router
    fmt.Println("Setting up Router...")
    engine := gin.Default()
    fmt.Println("Setting up Middlewares...")
    configCors := middlewares.CreateCorsConfig(config);
    engine.Use(cors.New(configCors))

    // Assemble API
    fmt.Println("Assembling Application...")
    programService := &controllers.ProgramService{ dbSql, dbSqlx }
    handler := router.Handler {
        engine,
        programService,
    }
    handler.SetupApiEndpoints()

    // Run web server
    handler.Engine.Run(":8080")
}
