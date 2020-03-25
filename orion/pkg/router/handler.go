package router

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine *gin.Engine
}

func (h *Handler) SetupApiEndpoints() {
	// h.Engine.Use(static.Serve("/", static.LocalFile("./sites/home", true)))
	h.Engine.Use(static.Serve("/", static.LocalFile("./sites/admin", true)))

	apiPrograms := h.Engine.Group("/api/programs/")
	{
		apiPrograms.GET("/v1/all", controllers.GetAllPrograms)
		apiPrograms.POST("/v1/create", controllers.CreateProgram)
		apiPrograms.GET("/v1/program/:programId", controllers.GetProgramById)
		apiPrograms.POST("/v1/program/:programId", controllers.UpdateProgram)
		apiPrograms.DELETE("/v1/program/:programId", controllers.DeleteProgram)
	}
	apiClasses := h.Engine.Group("api/classes/")
	{
		apiClasses.GET("/v1/class/:classId/sessions", controllers.GetAllSessionsByClassId)
	}
	apiLocations := h.Engine.Group("api/locations/")
	{
		apiLocations.GET("/v1/all", controllers.GetAllLocations)
		apiLocations.POST("/v1/create", controllers.CreateLocation)
		apiLocations.GET("/v1/location/:locId", controllers.GetLocationById)
		apiLocations.POST("/v1/location/:locId", controllers.UpdateLocation)
		apiLocations.DELETE("/v1/location/:locId", controllers.DeleteLocation)
	}
	apiAnnounces := h.Engine.Group("api/announcements/")
	{
		apiAnnounces.GET("/v1/all", controllers.GetAllAnnouncements)
		apiAnnounces.POST("/v1/create", controllers.CreateAnnouncement)
		apiAnnounces.GET("/v1/announcement/:id", controllers.GetAnnouncementById)
		apiAnnounces.POST("/v1/announcement/:id", controllers.UpdateAnnouncement)
		apiAnnounces.DELETE("/v1/announcement/:id", controllers.DeleteAnnouncement)
	}
	apiAchieves := h.Engine.Group("api/achievements/")
	{
		apiAchieves.GET("/v1/all", controllers.GetAllAchievements)
		apiAchieves.POST("/v1/create", controllers.CreateAchievement)
		apiAchieves.GET("/v1/achievement/:id", controllers.GetAchievementById)
		apiAchieves.POST("/v1/achievement/:id", controllers.UpdateAchievement)
		apiAchieves.DELETE("/v1/achievement/:id", controllers.DeleteAchievement)
	}
	apiSemesters := h.Engine.Group("api/semesters/")
	{
		apiSemesters.GET("/v1/all", controllers.GetAllSemesters)
		apiSemesters.POST("/v1/create", controllers.CreateSemester)
		apiSemesters.GET("/v1/semester/:semesterId", controllers.GetSemesterById)
		apiSemesters.POST("/v1/semester/:semesterId", controllers.UpdateSemester)
		apiSemesters.DELETE("/v1/semester/:semesterId", controllers.DeleteSemester)
	}
	apiSessions := h.Engine.Group("api/sessions/")
	{
		apiSessions.POST("/v1/create", controllers.CreateSession)
		apiSessions.GET("/v1/session/:id", controllers.GetSessionById)
		apiSessions.POST("/v1/session/:id", controllers.UpdateSession)
		apiSessions.DELETE("/v1/session/:id", controllers.DeleteSession)
	}
	// apiUsers := router.Group("api/users/")
	// apiAccounts := router.Group("api/accounts/")
}
