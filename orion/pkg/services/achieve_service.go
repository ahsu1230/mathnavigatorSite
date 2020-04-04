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
	GetAllGroupedByYear() ([]domains.AchieveYearGroup, error)
	GetUnpublished() ([]domains.Achieve, error)
	Create(domains.Achieve) error
	Update(uint, domains.Achieve) error
	Delete(uint) error
	Publish([]uint) error
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

func (as *achieveService) GetAllGroupedByYear() ([]domains.AchieveYearGroup, error) {
	achieves, err := repos.AchieveRepo.SelectAllGroupedByYear()
	if err != nil {
		return nil, err
	}
	return achieves, nil
}

func (as *achieveService) GetUnpublished() ([]domains.Achieve, error) {
	achieves, err := repos.AchieveRepo.SelectUnpublished()
	if err != nil {
		return nil, err
	}
	return achieves, nil
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

// TODO: Use DB Transactions
func (as *achieveService) Publish(ids []uint) error {
	for _, id := range ids {
		err := repos.AchieveRepo.Publish(id)
		if err != nil {
			return err
		}
	}
	return nil
}
