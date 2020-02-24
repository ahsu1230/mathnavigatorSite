package services

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/repos"
)

var ProgramService programServiceInterface = &programService{}

// Interface for ProgramService
type programServiceInterface interface {
    GetAll() ([]domains.Program, error)
    // GetByProgramId(c *gin.Context)
    // Create(c *gin.Context)
    // Update(c *gin.Context)
    // Delete(c *gin.Context)
}

// Struct that implements interface
type programService struct {}

func (ps *programService) GetAll() ([]domains.Program, error) {
	programs, err := repos.ProgramRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return programs, nil
}