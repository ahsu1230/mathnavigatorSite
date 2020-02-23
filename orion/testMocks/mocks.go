package testMocks

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/ahsu1230/mathnavigatorSite/orion/domains"
)

// ProgramService represents a mock implementation of domains.ProgramService
type ProgramService struct {
        GetAllFn      func(*gin.Context)
        GetAllInvoked bool
        GetByProgramIdFn      func(*gin.Context)
        GetByProgramIdInvoked bool
        CreateProgramFn      func(*gin.Context)
        CreateProgramInvoked bool
        UpdateProgramFn      func(*gin.Context)
        UpdateProgramInvoked bool
        DeleteProgramFn      func(*gin.Context)
        DeleteProgramInvoked bool
}

func (ps *ProgramService) GetAll(c *gin.Context) {
        ps.GetAllInvoked = true
        ps.GetAllFn(c)
}

func (ps *ProgramService) GetByProgramId(c *gin.Context) {
        ps.GetByProgramIdInvoked = true
        ps.GetByProgramIdFn(c)
}

func (ps *ProgramService) Create(c *gin.Context) {
        ps.CreateProgramInvoked = true
        ps.CreateProgramFn(c)
}

func (ps *ProgramService) Update(c *gin.Context) {
        ps.UpdateProgramInvoked = true
        ps.UpdateProgramFn(c)
}

func (ps *ProgramService) Delete(c *gin.Context) {
        ps.DeleteProgramInvoked = true
        ps.DeleteProgramFn(c)
}

func CreateMockProgram(rowId uint, programId string, programName string, grade1 uint, grade2 uint, description string) *domains.Program {
    return &domains.Program {
        rowId,
        1000,
        1000,
        sql.NullInt64{},
        programId,
        programName,
        grade1,
        grade2,
        description,
    }
}
