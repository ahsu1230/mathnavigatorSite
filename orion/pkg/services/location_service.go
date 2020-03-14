package services

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
)

var LocationService locationServiceInterface = &locationService{}

// Interface for LocationService
type locationServiceInterface interface {
	GetAll() ([]domains.Location, error)
	GetByLocationId(string) (domains.Location, error)
	Create(domains.Location) error
	Update(string, domains.Location) error
	Delete(string) error
}

// Struct that implements interface
type locationService struct{}

func (ps *locationService) GetAll() ([]domains.Location, error) {
	locations, err := repos.LocationRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (ps *locationService) GetByLocationId(locId string) (domains.Location, error) {
	location, err := repos.LocationRepo.SelectByLocationId(locId)
	if err != nil {
		return domains.Location{}, err
	}
	return location, nil
}

func (ps *locationService) Create(location domains.Location) error {
	err := repos.LocationRepo.Insert(location)
	return err
}

func (ps *locationService) Update(locId string, location domains.Location) error {
	err := repos.LocationRepo.Update(locId, location)
	return err
}

func (ps *locationService) Delete(locId string) error {
	err := repos.LocationRepo.Delete(locId)
	return err
}
