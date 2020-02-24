package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"

	"github.com/ahsu1230/mathnavigatorSite/orion/controllers/programs"
	"github.com/ahsu1230/mathnavigatorSite/orion/controllers/announce"
	"github.com/ahsu1230/mathnavigatorSite/orion/controllers/achieve"
	"github.com/ahsu1230/mathnavigatorSite/orion/controllers/semesters"
	"github.com/ahsu1230/mathnavigatorSite/orion/middlewares"
	"github.com/ahsu1230/mathnavigatorSite/orion/database"
)

func main() {
	fmt.Println("Orion service starting...")

	config := middlewares.RetrieveConfigurations()
	fmt.Println("Building server in mode: ", config.App.Build)

	fmt.Println("Connecting to DB...")
	configDb := config.Database
	database.OpenDb(configDb.Host, configDb.Port,
	configDb.Username, configDb.Password)
	fmt.Println("Performing DB Migrations...")
	database.Migrate()

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
	apiAnnounce := router.Group("api/announcements/")
	{
		apiAnnounce.GET("/v1/all", announce.GetAnnouncements)
		apiAnnounce.POST("/v1/create", announce.CreateAnnouncement)
		apiAnnounce.GET("/v1/announcement/:id", announce.GetAnnouncement)
		apiAnnounce.POST("/v1/announcement/:id", announce.UpdateAnnouncement)
		apiAnnounce.DELETE("/v1/announcement/:id", announce.DeleteAnnouncement)
	}
	apiAchieve := router.Group("api/achievements/")
	{
		apiAchieve.GET("/v1/all", achieve.GetAchievements)
		apiAchieve.POST("/v1/create", achieve.CreateAchievement)
		apiAchieve.GET("/v1/achievement/:id", achieve.GetAchievement)
		apiAchieve.POST("/v1/achievement/:id", achieve.UpdateAchievement)
		apiAchieve.DELETE("/v1/achievement/:id", achieve.DeleteAchievement)
	}
	apiSemesters := router.Group("api/semesters/")
	{
		apiSemesters.GET("/v1/all", semesters.GetSemesters)
		apiSemesters.POST("/v1/create", semesters.CreateSemester)
		apiSemesters.GET("/v1/semester/:semesterId", semesters.GetSemester)
		apiSemesters.POST("/v1/semester/:semesterId", semesters.UpdateSemester)
		apiSemesters.DELETE("/v1/semester/:semesterId", semesters.DeleteSemester)
	}
	// apiUsers := router.Group("api/users/")
	// apiAccounts := router.Group("api/accounts/")

	// Web server serves on :8080
	router.Run(":8080")

	// close DbSqlx when server finishes
	defer database.CloseDb()
}
