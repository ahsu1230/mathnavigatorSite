package services

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/repos"
)

var ProgramService programServiceInterface = &programService{}

// Interface for ProgramService
type programServiceInterface interface {
	GetAll(bool) ([]domains.Program, error)
	GetAllUnpublished() ([]domains.Program, error)
	GetByProgramId(string) (domains.Program, error)
	Create(domains.Program) error
	Update(string, domains.Program) error
	Publish([]string) error
	Delete(string) error
}

// Struct that implements interface
type programService struct{}

func (ps *programService) GetAll(publishedOnly bool) ([]domains.Program, error) {
	programs, err := repos.ProgramRepo.SelectAll(publishedOnly)
	if err != nil {
		return nil, err
	}
	return programs, nil
}

func (ps *programService) GetAllUnpublished() ([]domains.Program, error) {
	programs, err := repos.ProgramRepo.SelectAllUnpublished()
	if err != nil {
		return nil, err
	}
	return programs, nil
}

func (ps *programService) GetByProgramId(programId string) (domains.Program, error) {
	program, err := repos.ProgramRepo.SelectByProgramId(programId)
	if err != nil {
		return domains.Program{}, err
	}
	return program, nil
}

func (ps *programService) Create(program domains.Program) error {
	err := repos.ProgramRepo.Insert(program)
	return err
}

func (ps *programService) Update(programId string, program domains.Program) error {
	err := repos.ProgramRepo.Update(programId, program)
	return err
}

func (ps *programService) Publish(programIds []string) error {
	err := repos.ProgramRepo.Publish(programIds)
	return err
}

func (ps *programService) Delete(programId string) error {
	err := repos.ProgramRepo.Delete(programId)
	return err
}
