package main
import (
  "fmt"

  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/static"

  "orion/controllers/programs"
  "orion/middlewares"
  "orion/database"
)

func main() {
  fmt.Println("Orion service starting...")

  config := middlewares.RetrieveConfigurations()
  fmt.Println("Building server in mode: ", config.App.Build)

  fmt.Println("Connecting to DB...")
  configDb := config.Database
  database.OpenDb(configDb.Host, configDb.Port,
    configDb.Username, configDb.Password)

  fmt.Println("Setting up Router...")
  router := gin.Default()

  fmt.Println("Setting up Middlewares...")

  // CORS middleware
  configCors := middlewares.CreateCorsConfig(config);
  router.Use(cors.New(configCors))

  // Webpage Routers
  // router.Use(static.Serve("/", static.LocalFile("./sites/home", true)))
  router.Use(static.Serve("/", static.LocalFile("./sites/admin", true)))

  // API Routers
  apiPrograms := router.Group("/api/programs/")
  {
    apiPrograms.GET("/v1/all", programs.GetPrograms)
    apiPrograms.POST("/v1/create", programs.CreateProgram)
    apiPrograms.GET("/v1/program/:programId", programs.GetProgram)
    apiPrograms.POST("/v1/program/:programId", programs.UpdateProgram)
    apiPrograms.DELETE("/v1/program/:programId", programs.DeleteProgram)
  }
  // apiClasses := router.Group("api/classes/")
  // apiLocations := router.Group("api/locations/")
  // apiAnnounce := router.Group("api/announce/")
  // apiAchieve := router.Group("api/achieve/")
  // apiSemesters := router.Group("api/semesters/")
  // apiUsers := router.Group("api/users/")
  // apiAccounts := router.Group("api/accounts/")

  // Web server serves on :8080
	router.Run(":8080")

  // close DbConn when server finishes
  defer database.CloseDb();
}
