package services

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
)

var UserService userServiceInterface = &userService{}

// Interface for UserService
type userServiceInterface interface {
	GetAll() ([]domains.User, error)
	GetById(uint) (domains.User, error)
	Create(domains.User) error
	Update(uint, domains.User) error
	Delete(uint) error
}

// Struct that implements interface
type userService struct{}

func (us *userService) GetAll() ([]domains.User, error) {
	users, err := repos.UserRepo.SelectAll()
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
