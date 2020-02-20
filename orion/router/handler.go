package router

import (
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
    "github.com/ahsu1230/mathnavigatorSite/orion/domains"
)

type Handler struct {
    Engine *gin.Engine
    ProgramService domains.ProgramService
}

func (h *Handler) SetupApiEndpoints() {
    // h.Engine.Use(static.Serve("/", static.LocalFile("./sites/home", true)))
    h.Engine.Use(static.Serve("/", static.LocalFile("./sites/admin", true)))

    apiPrograms := h.Engine.Group("/api/programs/")
    {
        apiPrograms.GET("/v1/all", h.ProgramService.GetAll)
        apiPrograms.POST("/v1/create", h.ProgramService.Create)
        apiPrograms.GET("/v1/program/:programId", h.ProgramService.GetByProgramId)
        apiPrograms.POST("/v1/program/:programId", h.ProgramService.Update)
        apiPrograms.DELETE("/v1/program/:programId", h.ProgramService.Delete)
    }
    // apiClasses := router.Group("api/classes/")
    // apiLocations := router.Group("api/locations/")
    // apiAnnounce := router.Group("api/announce/")
    // apiAchieve := router.Group("api/achieve/")
    // apiSemesters := router.Group("api/semesters/")
    // apiUsers := router.Group("api/users/")
    // apiAccounts := router.Group("api/accounts/")
}
