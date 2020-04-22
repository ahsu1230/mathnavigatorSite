package services

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/repos"
)

var LocationService locationServiceInterface = &locationService{}

// Interface for LocationService
type locationServiceInterface interface {
	GetAll(bool) ([]domains.Location, error)
	GetAllUnpublished() ([]domains.Location, error)
	GetByLocationId(string) (domains.Location, error)
	Create(domains.Location) error
	Update(string, domains.Location) error
	Publish([]string) error
	Delete(string) error
}

// Struct that implements interface
type locationService struct{}

func (ls *locationService) GetAll(publishedOnly bool) ([]domains.Location, error) {
	locations, err := repos.LocationRepo.SelectAll(publishedOnly)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (ls *locationService) GetAllUnpublished() ([]domains.Location, error) {
	locations, err := repos.LocationRepo.SelectAllUnpublished()
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (ls *locationService) GetByLocationId(locationId string) (domains.Location, error) {
	location, err := repos.LocationRepo.SelectByLocationId(locationId)
	if err != nil {
		return domains.Location{}, err
	}
	return location, nil
}

func (ls *locationService) Create(location domains.Location) error {
	err := repos.LocationRepo.Insert(location)
	return err
}
func (ls *locationService) Update(locationId string, location domains.Location) error {
	err := repos.LocationRepo.Update(locationId, location)
	return err
}
func (ls *locationService) Publish(locationIds []string) error {
	err := repos.LocationRepo.Publish(locationIds)
	return err
}
func (ls *locationService) Delete(locationId string) error {
	err := repos.LocationRepo.Delete(locationId)
	return err
}
