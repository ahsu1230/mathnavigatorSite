package services

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
)

var ClassService classServiceInterface = &classService{}

// Interface for ClassService
type classServiceInterface interface {
	GetAll() ([]domains.Class, error)
	GetByClassId(string) (domains.Class, error)
	Create(domains.Class) error
	Update(string, domains.Class) error
	Delete(string) error
}

// Struct that implements interface
type classService struct{}

func (cs *classService) GetAll() ([]domains.Class, error) {
	classes, err := repos.ClassRepo.SelectAll()
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

func (cs *classService) Create(class domains.Class) error {
	err := repos.ClassRepo.Insert(class)
	return err
}

func (cs *classService) Update(classId string, class domains.Class) error {
	err := repos.ClassRepo.Update(classId, class)
	return err
}

func (cs *classService) Delete(classId string) error {
	err := repos.ClassRepo.Delete(classId)
	return err
}
