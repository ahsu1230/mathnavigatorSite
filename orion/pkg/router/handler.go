package router

import (
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
    "github.com/ahsu1230/mathnavigatorSite/orion/pkg/controllers"
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
    // apiAnnounce := router.Group("api/announce/")
    // apiAchieve := router.Group("api/achieve/")
    // apiSemesters := router.Group("api/semesters/")
    // apiUsers := router.Group("api/users/")
    // apiAccounts := router.Group("api/accounts/")
}
