package router

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine *gin.Engine
}

func (h *Handler) SetupApiEndpoints() {
	// h.Engine.Use(static.Serve("/", static.LocalFile("./sites/home", true)))
	h.Engine.Use(static.Serve("/", static.LocalFile("./sites/admin", true)))

	h.Engine.GET("/api/unpublished", controllers.GetAllUnpublished)
	apiPrograms := h.Engine.Group("/api/programs")
	{
		apiPrograms.GET("/all", controllers.GetAllPrograms)
		apiPrograms.POST("/create", controllers.CreateProgram)
		apiPrograms.POST("/publish", controllers.PublishPrograms)
		apiPrograms.GET("/program/:programId", controllers.GetProgramById)
		apiPrograms.POST("/program/:programId", controllers.UpdateProgram)
		apiPrograms.DELETE("/program/:programId", controllers.DeleteProgram)
	}
	apiClasses := h.Engine.Group("api/classes")
	{
		apiClasses.GET("/all", controllers.GetAllClasses)
		apiClasses.POST("/create", controllers.CreateClass)
		apiClasses.POST("/publish", controllers.PublishClasses)
		apiClasses.GET("/class/:classId", controllers.GetClassById)
		apiClasses.POST("/class/:classId", controllers.UpdateClass)
		apiClasses.DELETE("/class/:classId", controllers.DeleteClass)
		apiClasses.GET("/classes/program/:programId", controllers.GetClassesByProgram)
		apiClasses.GET("/classes/semester/:semesterId", controllers.GetClassesBySemester)
		apiClasses.GET("/classes/program/:programId/semester/:semesterId", controllers.GetClassesByProgramAndSemester)
	}
	apiLocations := h.Engine.Group("api/locations")
	{
		apiLocations.GET("/all", controllers.GetAllLocations)
		apiLocations.POST("/create", controllers.CreateLocation)
		apiLocations.POST("/publish", controllers.PublishLocations)
		apiLocations.GET("/location/:locationId", controllers.GetLocationById)
		apiLocations.POST("/location/:locationId", controllers.UpdateLocation)
		apiLocations.DELETE("/location/:locationId", controllers.DeleteLocation)
	}
	apiAnnounces := h.Engine.Group("api/announcements")
	{
		apiAnnounces.GET("/all", controllers.GetAllAnnouncements)
		apiAnnounces.POST("/create", controllers.CreateAnnouncement)
		apiAnnounces.GET("/announcement/:id", controllers.GetAnnouncementById)
		apiAnnounces.POST("/announcement/:id", controllers.UpdateAnnouncement)
		apiAnnounces.DELETE("/announcement/:id", controllers.DeleteAnnouncement)
	}
	apiAchieves := h.Engine.Group("api/achievements")
	{
		apiAchieves.GET("/all", controllers.GetAllAchievements)
		apiAchieves.GET("/years", controllers.GetAllAchievementsGroupedByYear)
		apiAchieves.POST("/create", controllers.CreateAchievement)
		apiAchieves.POST("/publish", controllers.PublishAchievements)
		apiAchieves.GET("/achievement/:id", controllers.GetAchievementById)
		apiAchieves.POST("/achievement/:id", controllers.UpdateAchievement)
		apiAchieves.DELETE("/achievement/:id", controllers.DeleteAchievement)
	}
	apiSemesters := h.Engine.Group("api/semesters")
	{
		apiSemesters.GET("/all", controllers.GetAllSemesters)
		apiSemesters.POST("/create", controllers.CreateSemester)
		apiSemesters.POST("/publish", controllers.PublishSemesters)
		apiSemesters.GET("/semester/:semesterId", controllers.GetSemesterById)
		apiSemesters.POST("/semester/:semesterId", controllers.UpdateSemester)
		apiSemesters.DELETE("/semester/:semesterId", controllers.DeleteSemester)
	}
	apiSessions := h.Engine.Group("api/sessions")
	{
		apiSessions.POST("/create", controllers.CreateSessions)
		apiSessions.POST("/publish", controllers.PublishSessions)
		apiSessions.GET("/session/:id", controllers.GetSessionById)
		apiSessions.POST("/session/:id", controllers.UpdateSession)
		apiSessions.DELETE("/delete", controllers.DeleteSessions)
		apiSessions.GET("/class/:classId", controllers.GetAllSessionsByClassId)
	}
	apiUsers := h.Engine.Group("api/users")
	{
		apiUsers.POST("/create", controllers.CreateUser)
		apiUsers.GET("/user/:id", controllers.GetUserById)
		apiUsers.POST("/user/:id", controllers.UpdateUser)
		apiUsers.DELETE("/user/:id", controllers.DeleteUser)
		apiUsers.GET("/account/:accountId", controllers.GetUsersByAccountId)
		apiUsers.POST("/search", controllers.SearchUsers)
	}

	apiAccounts := h.Engine.Group("api/accounts")
	{
		apiAccounts.POST("/create", controllers.CreateAccount)
		apiAccounts.GET("/account/:id", controllers.GetAccountById)
		apiAccounts.POST("/account/:id", controllers.UpdateAccount)
		apiAccounts.DELETE("/account/:id", controllers.DeleteAccount)
		apiAccounts.POST("/search", controllers.SearchAccount)
	}

	apiAFH := h.Engine.Group("api/askforhelp")
	{
		apiAFH.GET("/all", controllers.GetAllAFH)
		apiAFH.POST("/create", controllers.CreateAFH)
		apiAFH.GET("/afh/:id", controllers.GetAFHById)
		apiAFH.POST("/afh/:id", controllers.UpdateAFH)
		apiAFH.DELETE("/afh/:id", controllers.DeleteAFH)
	}

	apiTransaction := h.Engine.Group("api/transactions")
	{
		apiTransaction.GET("/all", controllers.GetAllTransactions)
		apiTransaction.POST("/create", controllers.CreateTransaction)
		apiTransaction.GET("/transaction/:id", controllers.GetTransactionById)
		apiTransaction.POST("/transaction/:id", controllers.UpdateTransaction)
		apiTransaction.DELETE("/transaction/:id", controllers.DeleteTransaction)
	}
	// apiAccounts := router.Group("api/accounts/")
}
