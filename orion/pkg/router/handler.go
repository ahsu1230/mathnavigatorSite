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
	// apiClasses := router.Group("api/classes/")
	// apiLocations := router.Group("api/locations/")
	apiAnnounces := h.Engine.Group("api/announcements/")
	{
		apiAnnounces.GET("/v1/all", controllers.GetAllAnnouncements)
		apiAnnounces.POST("/v1/create", controllers.CreateAnnouncement)
		apiAnnounces.GET("/v1/announce/:id", controllers.GetAnnouncementById)
		apiAnnounces.POST("/v1/announce/:id", controllers.UpdateAnnouncement)
		apiAnnounces.DELETE("/v1/announce/:id", controllers.DeleteAnnouncement)
	}
	// apiAchieve := router.Group("api/achieve/")
	apiSemesters := h.Engine.Group("api/semesters/")
	{
		apiSemesters.GET("/v1/all", controllers.GetAllSemesters)
		apiSemesters.POST("/v1/create", controllers.CreateSemester)
		apiSemesters.GET("/v1/semester/:id", controllers.GetSemesterById)
		apiSemesters.POST("/v1/semester/:id", controllers.UpdateSemester)
		apiSemesters.DELETE("/v1/semester/:id", controllers.DeleteSemester)
	}
	// apiUsers := router.Group("api/users/")
	// apiAccounts := router.Group("api/accounts/")
}
