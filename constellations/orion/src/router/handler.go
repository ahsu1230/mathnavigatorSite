package router

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine *gin.Engine
}

func (h *Handler) SetupApiEndpoints() {
	h.Engine.GET("/api/classesbysemesters", controllers.GetAllProgramsSemestersClasses)
	apiPrograms := h.Engine.Group("/api/programs")
	{
		apiPrograms.GET("/all", controllers.GetAllPrograms)
		apiPrograms.GET("/featured", controllers.GetAllProgramFeatured)
		apiPrograms.GET("/program/:programId", controllers.GetProgramById)
		apiPrograms.POST("/create", controllers.CreateProgram)
		apiPrograms.POST("/program/:programId", controllers.UpdateProgram)
		apiPrograms.DELETE("/program/:programId", controllers.DeleteProgram)
	}
	apiClasses := h.Engine.Group("api/classes")
	{
		apiClasses.GET("/all", controllers.GetAllClasses)
		apiClasses.GET("/full-states", controllers.GetFullStates)
		apiClasses.GET("/class/:classId", controllers.GetClassById)
		apiClasses.GET("/semester/:semesterId", controllers.GetClassesBySemester)
		apiClasses.GET("/program/:programId", controllers.GetClassesByProgram)
		apiClasses.GET("/program/:programId/semester/:semesterId", controllers.GetClassesByProgramAndSemester)
		apiClasses.GET("/unpublished", controllers.GetUnpublishedClasses)
		apiClasses.POST("/create", controllers.CreateClass)
		apiClasses.POST("/publish", controllers.PublishClasses)
		apiClasses.POST("/class/:classId", controllers.UpdateClass)
		apiClasses.DELETE("/class/:classId", controllers.DeleteClass)
	}
	apiLocations := h.Engine.Group("api/locations")
	{
		apiLocations.GET("/all", controllers.GetAllLocations)
		apiLocations.GET("/location/:locationId", controllers.GetLocationById)
		apiLocations.POST("/create", controllers.CreateLocation)
		apiLocations.POST("/location/:locationId", controllers.UpdateLocation)
		apiLocations.DELETE("/location/:locationId", controllers.DeleteLocation)
	}
	apiAnnounces := h.Engine.Group("api/announcements")
	{
		apiAnnounces.GET("/all", controllers.GetAllAnnouncements)
		apiAnnounces.GET("/announcement/:id", controllers.GetAnnouncementById)
		apiAnnounces.POST("/create", controllers.CreateAnnouncement)
		apiAnnounces.POST("/announcement/:id", controllers.UpdateAnnouncement)
		apiAnnounces.DELETE("/announcement/:id", controllers.DeleteAnnouncement)
	}
	apiAchieves := h.Engine.Group("api/achievements")
	{
		apiAchieves.GET("/all", controllers.GetAllAchievements)
		apiAchieves.GET("/years", controllers.GetAllAchievementsGroupedByYear)
		apiAchieves.GET("/achievement/:id", controllers.GetAchievementById)
		apiAchieves.POST("/create", controllers.CreateAchievement)
		apiAchieves.POST("/achievement/:id", controllers.UpdateAchievement)
		apiAchieves.DELETE("/achievement/:id", controllers.DeleteAchievement)
	}
	apiSemesters := h.Engine.Group("api/semesters")
	{
		apiSemesters.GET("/all", controllers.GetAllSemesters)
		apiSemesters.GET("/seasons", controllers.GetAllSeasons)
		apiSemesters.GET("/semester/:semesterId", controllers.GetSemesterById)
		apiSemesters.POST("/create", controllers.CreateSemester)
		apiSemesters.POST("/semester/:semesterId", controllers.UpdateSemester)
		apiSemesters.DELETE("/semester/:semesterId", controllers.DeleteSemester)
	}
	apiSessions := h.Engine.Group("api/sessions")
	{
		apiSessions.GET("/class/:classId", controllers.GetAllSessionsByClassId)
		apiSessions.GET("/session/:id", controllers.GetSessionById)
		apiSessions.POST("/create", controllers.CreateSessions)
		apiSessions.POST("/session/:id", controllers.UpdateSession)
		apiSessions.DELETE("/delete", controllers.DeleteSessions)
	}
	apiAFH := h.Engine.Group("api/askforhelp")
	{
		apiAFH.GET("/all", controllers.GetAllAFH)
		apiAFH.GET("/subjects", controllers.GetAllAFHSubjects)
		apiAFH.GET("/afh/:id", controllers.GetAFHById)
		apiAFH.POST("/create", controllers.CreateAFH)
		apiAFH.POST("/afh/:id", controllers.UpdateAFH)
		apiAFH.DELETE("/afh/:id", controllers.DeleteAFH)
	}

	apiAccounts := h.Engine.Group("api/accounts")
	{
		apiAccounts.GET("/account/:id", controllers.GetAccountById)
		apiAccounts.GET("/unpaid", controllers.GetNegativeBalanceAccounts)
		apiAccounts.POST("/create", controllers.CreateAccountAndUser)
		apiAccounts.POST("/account/:id", controllers.UpdateAccount)
		apiAccounts.DELETE("/account/:id", controllers.DeleteAccount)
		apiAccounts.POST("/search", controllers.SearchAccount)
	}
	apiTransaction := h.Engine.Group("api/transactions")
	{
		apiTransaction.GET("/account/:accountId", controllers.GetTransactionsByAccountId)
		apiTransaction.GET("/transaction/:id", controllers.GetTransactionById)
		apiTransaction.GET("/types", controllers.GetAllPaymentTypes)
		apiTransaction.POST("/create", controllers.CreateTransaction)
		apiTransaction.POST("/transaction/:id", controllers.UpdateTransaction)
		apiTransaction.DELETE("/transaction/:id", controllers.DeleteTransaction)
	}
	apiUsers := h.Engine.Group("api/users")
	{
		apiUsers.GET("/new", controllers.GetNewUsers)
		apiUsers.GET("/account/:accountId", controllers.GetUsersByAccountId)
		apiUsers.GET("/user/:id", controllers.GetUserById)
		apiUsers.POST("/search", controllers.SearchUsers)
		apiUsers.POST("/create", controllers.CreateUser)
		apiUsers.POST("/user/:id", controllers.UpdateUser)
		apiUsers.DELETE("/user/:id", controllers.DeleteUser)
	}
	apiUserClasses := h.Engine.Group("api/user-classes")
	{
		apiUserClasses.GET("/class/:classId", controllers.GetUsersByClassId)
		apiUserClasses.GET("/class/:classId/user/:userId", controllers.GetUserClassByUserAndClass)
		apiUserClasses.GET("/user/:userId", controllers.GetClassesByUserId)
		apiUserClasses.GET("/new", controllers.GetNewClasses)
		apiUserClasses.GET("/states", controllers.GetStateValues)
		apiUserClasses.POST("/create", controllers.CreateUserClass)
		apiUserClasses.POST("/user-class/:id", controllers.UpdateUserClass)
		apiUserClasses.DELETE("/user-class/:id", controllers.DeleteUserClass)
	}
	apiUserAfh := h.Engine.Group("api/user-afhs")
	{
		apiUserAfh.GET("/new", controllers.GetUserAfhByNew)
		apiUserAfh.GET("afh/:afhId", controllers.GetUserAfhByAfhId)
		apiUserAfh.GET("/users/:userId/afh/:afhId", controllers.GetUserAfhByBothIds)
		apiUserAfh.GET("/users/:userId", controllers.GetUserAfhByUserId)
		apiUserAfh.POST("/create", controllers.CreateUserAfh)
		apiUserAfh.POST("/user-afh/:id", controllers.UpdateUserAfh)
		apiUserAfh.DELETE("/user-afh/:id", controllers.DeleteUserAfh)
	}
}
