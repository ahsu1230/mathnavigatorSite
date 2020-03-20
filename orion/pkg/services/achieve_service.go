package services

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
)

var AchieveService achieveServiceInterface = &achieveService{}

// Interface for AchieveService
type achieveServiceInterface interface {
	GetAll() ([]domains.Achieve, error)
	GetById(uint) (domains.Achieve, error)
	Create(domains.Achieve) error
	Update(uint, domains.Achieve) error
	Delete(uint) error
}

// Struct that implements interface
type achieveService struct{}

func (as *achieveService) GetAll() ([]domains.Achieve, error) {
	achieves, err := repos.AchieveRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return achieves, nil
}

func (as *achieveService) GetById(id uint) (domains.Achieve, error) {
	achieve, err := repos.AchieveRepo.SelectById(id)
	if err != nil {
		return domains.Achieve{}, err
	}
	return achieve, nil
}

func (as *achieveService) Create(achieve domains.Achieve) error {
	err := repos.AchieveRepo.Insert(achieve)
	return err
}

func (as *achieveService) Update(id uint, achieve domains.Achieve) error {
	err := repos.AchieveRepo.Update(id, achieve)
	return err
}

func (as *achieveService) Delete(id uint) error {
	err := repos.AchieveRepo.Delete(id)
	return err
}