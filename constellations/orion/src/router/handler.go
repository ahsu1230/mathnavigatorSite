package router

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine *gin.Engine
}

func (h *Handler) SetupApiEndpoints() {
	apiPrograms := h.Engine.Group("/api/programs")
	{
		apiPrograms.GET("/all", controllers.GetAllPrograms)
		apiPrograms.GET("/states", controllers.GetAllProgramStates)
		apiPrograms.GET("/program/:programId", controllers.GetProgramById)
		apiPrograms.POST("/create", controllers.CreateProgram)
		apiPrograms.POST("/program/:programId", controllers.UpdateProgram)
		apiPrograms.DELETE("/program/:programId", controllers.DeleteProgram)
	}
	apiClasses := h.Engine.Group("api/classes")
	{
		apiClasses.GET("/all", controllers.GetAllClasses)
		apiClasses.GET("/class/:classId", controllers.GetClassById)
		apiClasses.GET("/classes/program/:programId", controllers.GetClassesByProgram)
		apiClasses.GET("/classes/semester/:semesterId", controllers.GetClassesBySemester)
		apiClasses.GET("/classes/program/:programId/semester/:semesterId", controllers.GetClassesByProgramAndSemester)
		apiClasses.GET("/unpublished", controllers.GetUnpublishedClasses)
		apiClasses.POST("/create", controllers.CreateClass)
		apiClasses.POST("/publish", controllers.PublishClasses)
		apiClasses.POST("/class/:classId", controllers.UpdateClass)
		apiClasses.DELETE("/class/:classId", controllers.DeleteClass)
	}
	apiLocations := h.Engine.Group("api/locations")
	{
		apiLocations.GET("/all", controllers.GetAllLocations)
		apiLocations.POST("/create", controllers.CreateLocation)
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
		apiAchieves.GET("/achievement/:id", controllers.GetAchievementById)
		apiAchieves.POST("/achievement/:id", controllers.UpdateAchievement)
		apiAchieves.DELETE("/achievement/:id", controllers.DeleteAchievement)
	}
	apiSemesters := h.Engine.Group("api/semesters")
	{
		apiSemesters.GET("/all", controllers.GetAllSemesters)
		apiSemesters.GET("/seasons", controllers.GetAllSeasons)
		apiSemesters.POST("/create", controllers.CreateSemester)
		apiSemesters.GET("/semester/:semesterId", controllers.GetSemesterById)
		apiSemesters.POST("/semester/:semesterId", controllers.UpdateSemester)
		apiSemesters.DELETE("/semester/:semesterId", controllers.DeleteSemester)
	}
	apiSessions := h.Engine.Group("api/sessions")
	{
		apiSessions.POST("/create", controllers.CreateSessions)
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
		apiUsers.GET("/new", controllers.GetNewUsers)
	}

	apiAccounts := h.Engine.Group("api/accounts")
	{
		apiAccounts.POST("/create", controllers.CreateAccountAndUser)
		apiAccounts.GET("/account/:id", controllers.GetAccountById)
		apiAccounts.GET("/unpaid", controllers.GetNegativeBalanceAccounts)
		apiAccounts.POST("/account/:id", controllers.UpdateAccount)
		apiAccounts.DELETE("/account/:id", controllers.DeleteAccount)
		apiAccounts.POST("/search", controllers.SearchAccount)
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

	apiUserClasses := h.Engine.Group("api/user-classes")
	{
		apiUserClasses.POST("/create", controllers.CreateUserClass)
		apiUserClasses.GET("/class/:classId", controllers.GetUsersByClassId)
		apiUserClasses.GET("/user/:userId", controllers.GetClassesByUserId)
		apiUserClasses.GET("/class/:classId/user/:userId", controllers.GetUserClassByUserAndClass)
		apiUserClasses.POST("/user-class/:id", controllers.UpdateUserClass)
		apiUserClasses.DELETE("/user-class/:id", controllers.DeleteUserClass)
		apiUserClasses.GET("/states", controllers.GetStateValues)
		apiUserClasses.GET("/new", controllers.GetNewClasses)
	}
	apiTransaction := h.Engine.Group("api/transactions")
	{
		apiTransaction.GET("/account/:accountId", controllers.GetTransactionsByAccountId)
		apiTransaction.POST("/create", controllers.CreateTransaction)
		apiTransaction.GET("/transaction/:id", controllers.GetTransactionById)
		apiTransaction.POST("/transaction/:id", controllers.UpdateTransaction)
		apiTransaction.DELETE("/transaction/:id", controllers.DeleteTransaction)
		apiTransaction.GET("/types", controllers.GetAllPaymentTypes)
	}

	apiUserAfh := h.Engine.Group("api/userafhs")
	{
		apiUserAfh.GET("/users/:userId", controllers.GetUserAfhByUserId)
		apiUserAfh.GET("afh/:afhId", controllers.GetUserAfhByAfhId)
		apiUserAfh.GET("/users/:userId/afh/:afhId", controllers.GetUserAfhByBothIds)
		apiUserAfh.GET("/new", controllers.GetUserAfhByNew)
		apiUserAfh.POST("/create", controllers.CreateUserAfh)
		apiUserAfh.POST("/userafh/:id", controllers.UpdateUserAfh)
		apiUserAfh.DELETE("/userafh/:id", controllers.DeleteUserAfh)
	}

	h.Engine.GET("/api/classesbysemesters", controllers.GetAllProgramsSemestersClasses)
}
