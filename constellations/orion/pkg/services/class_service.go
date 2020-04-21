package services

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/repos"
)

var ClassService classServiceInterface = &classService{}

// Interface for ClassService
type classServiceInterface interface {
	GetAll(bool) ([]domains.Class, error)
	GetAllUnpublished() ([]domains.Class, error)
	GetByClassId(string) (domains.Class, error)
	GetByProgramId(string) ([]domains.Class, error)
	GetBySemesterId(string) ([]domains.Class, error)
	GetByProgramAndSemesterId(string, string) ([]domains.Class, error)
	Create(domains.Class) error
	Update(string, domains.Class) error
	Publish([]string) error
	Delete(string) error
}

// Struct that implements interface
type classService struct{}

func (cs *classService) GetAll(publishedOnly bool) ([]domains.Class, error) {
	classes, err := repos.ClassRepo.SelectAll(publishedOnly)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (cs *classService) GetAllUnpublished() ([]domains.Class, error) {
	classes, err := repos.ClassRepo.SelectAllUnpublished()
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (cs *classService) GetByClassId(classId string) (domains.Class, error) {
	class, err := repos.ClassRepo.SelectByClassId(classId)
	if err != nil {
		return domains.Class{}, err
	}
	return class, nil
}

func (cs *classService) GetByProgramId(programId string) ([]domains.Class, error) {
	classes, err := repos.ClassRepo.SelectByProgramId(programId)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (cs *classService) GetBySemesterId(semesterId string) ([]domains.Class, error) {
	classes, err := repos.ClassRepo.SelectBySemesterId(semesterId)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (cs *classService) GetByProgramAndSemesterId(programId, semesterId string) ([]domains.Class, error) {
	classes, err := repos.ClassRepo.SelectByProgramAndSemesterId(programId, semesterId)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (cs *classService) Create(class domains.Class) error {
	err := repos.ClassRepo.Insert(class)
	return err
}

func (cs *classService) Update(classId string, class domains.Class) error {
	err := repos.ClassRepo.Update(classId, class)
	return err
}

func (cs *classService) Publish(classIds []string) error {
	err := repos.ClassRepo.Publish(classIds)
	return err
}

func (cs *classService) Delete(classId string) error {
	err := repos.ClassRepo.Delete(classId)
	return err
}
