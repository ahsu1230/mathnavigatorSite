package main
import (
  "fmt"
  // "net/http"

  "github.com/gin-gonic/gin"
  "orion/controllers"
)

func main() {
  fmt.Println("Orion service starting...")

  router := gin.Default()

	// router.GET("/", controllers.SayHello)
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
}
