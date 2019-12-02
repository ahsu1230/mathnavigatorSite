package main
import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/static"

  "orion/controllers"
  "orion/models"
)

func main() {
  fmt.Println("Orion service starting...")

  fmt.Println("Connecting to DB...")
  models.OpenDb()

  fmt.Println("Setting up Router...")
  router := gin.Default()

  // Webpage Routers
  router.Use(static.Serve("/", static.LocalFile("./sites/home", true)))
  router.Use(static.Serve("/admin", static.LocalFile("./sites/admin", true)))

  // API Routers
  apiPrograms := router.Group("/api/programs/")
  {
    apiPrograms.GET("/v1", controllers.GetPrograms)
    apiPrograms.POST("/v1/create", controllers.CreateProgram)
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
