package main
import (
  "fmt"
  "os"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/static"
  "gopkg.in/yaml.v2"

  "orion/controllers"
  "orion/models"
)

type Config struct {
    App struct {
      Build string `yaml:"build"`
    } `yaml:"app"`
    Database struct {
      Host string `yaml:"host"`
      Port int `yaml:"port"`
      Username string `yaml:"user"`
      Password string `yaml:"pass"`
    } `yaml:"database"`
}

func retrieveConfigurations() (Config) {
  configFile := os.Args[1]
  fmt.Println("Configuration File: ", configFile)

  file, errFile := os.Open(configFile)
  if errFile != nil {
    fmt.Println("Error with file ", configFile)
  }

  var cfg Config
  decoder := yaml.NewDecoder(file)
  errParse := decoder.Decode(&cfg)
  if errParse != nil {
      fmt.Println("Error from parsing ", errParse)
  }
  return cfg
}

func main() {
  fmt.Println("Orion service starting...")

  config := retrieveConfigurations()
  fmt.Println("Building server in mode: ", config.App.Build)

  fmt.Println("Connecting to DB...")
  db := config.Database
  models.OpenDb(db.Host, db.Port, db.Username, db.Password)

  fmt.Println("Setting up Router...")
  router := gin.Default()

  // Webpage Routers
  // router.Use(static.Serve("/", static.LocalFile("./sites/home", true)))
  router.Use(static.Serve("/", static.LocalFile("./sites/admin", true)))

  // API Routers
  apiPrograms := router.Group("/api/programs/")
  {
    apiPrograms.GET("/v1", controllers.GetPrograms)
    apiPrograms.POST("/v1/create", controllers.CreateProgram)
    apiPrograms.GET("/v1/program/:programId", controllers.GetProgram)
    apiPrograms.POST("/v1/program/:programId", controllers.UpdateProgram)
    apiPrograms.DELETE("/v1/program/:programId", controllers.DeleteProgram)
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

  // Close db?
}
