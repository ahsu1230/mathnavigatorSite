package services

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
)

var UserService userServiceInterface = &userService{}

// Interface for UserService
type userServiceInterface interface {
	GetAll(string, int, int) ([]domains.User, error)
	GetById(uint) (domains.User, error)
	GetByGuardianId(uint) ([]domains.User, error)
	Create(domains.User) error
	Update(uint, domains.User) error
	Delete(uint) error
}

// Struct that implements interface
type userService struct{}

func (us *userService) GetAll(search string, pageSize, offset int) ([]domains.User, error) {
	users, err := repos.UserRepo.SelectAll(search, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userService) GetById(id uint) (domains.User, error) {
	user, err := repos.UserRepo.SelectById(id)
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (us *userService) GetByGuardianId(guardianId uint) ([]domains.User, error) {
	user, err := repos.UserRepo.SelectByGuardianId(guardianId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) Create(user domains.User) error {
	err := repos.UserRepo.Insert(user)
	return err
}

func (us *userService) Update(id uint, user domains.User) error {
	err := repos.UserRepo.Update(id, user)
	return err
}

func (us *userService) Delete(id uint) error {
	err := repos.UserRepo.Delete(id)
	return err
}
