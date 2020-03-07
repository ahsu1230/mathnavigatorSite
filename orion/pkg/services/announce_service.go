package services

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
)

var AnnounceService announceServiceInterface = &announceService{}

// Interface for AnnounceService
type announceServiceInterface interface {
	GetAll() ([]domains.Announce, error)
	GetByAnnounceId(uint) (domains.Announce, error)
	Create(domains.Announce) error
	Update(uint, domains.Announce) error
	Delete(uint) error
}

// Struct that implements interface
type announceService struct{}

func (as *announceService) GetAll() ([]domains.Announce, error) {
	announces, err := repos.AnnounceRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return announces, nil
}

func (as *announceService) GetByAnnounceId(id uint) (domains.Announce, error) {
	announce, err := repos.AnnounceRepo.SelectByAnnounceId(id)
	if err != nil {
		return domains.Announce{}, err
	}
	return announce, nil
}

func (as *announceService) Create(announce domains.Announce) error {
	err := repos.AnnounceRepo.Insert(announce)
	return err
}

func (as *announceService) Update(id uint, announce domains.Announce) error {
	err := repos.AnnounceRepo.Update(id, announce)
	return err
}

func (as *announceService) Delete(id uint) error {
	err := repos.AnnounceRepo.Delete(id)
	return err
}
