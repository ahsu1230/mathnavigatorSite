package services

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/repos"
)

var SemesterService semesterServiceInterface = &semesterService{}

// Interface for SemesterService
type semesterServiceInterface interface {
	GetAll(bool) ([]domains.Semester, error)
	GetAllUnpublished() ([]domains.Semester, error)
	GetBySemesterId(string) (domains.Semester, error)
	Create(domains.Semester) error
	Update(string, domains.Semester) error
	Publish([]string) error
	Delete(string) error
}

// Struct that implements interface
type semesterService struct{}

func (ss *semesterService) GetAll(publishedOnly bool) ([]domains.Semester, error) {
	semesters, err := repos.SemesterRepo.SelectAll(publishedOnly)
	if err != nil {
		return nil, err
	}
	return semesters, nil
}

func (ss *semesterService) GetAllUnpublished() ([]domains.Semester, error) {
	semesters, err := repos.SemesterRepo.SelectAllUnpublished()
	if err != nil {
		return nil, err
	}
	return semesters, nil
}

func (ss *semesterService) GetBySemesterId(semesterId string) (domains.Semester, error) {
	semester, err := repos.SemesterRepo.SelectBySemesterId(semesterId)
	if err != nil {
		return domains.Semester{}, err
	}
	return semester, nil
}

func (ss *semesterService) Create(semester domains.Semester) error {
	err := repos.SemesterRepo.Insert(semester)
	return err
}

func (ss *semesterService) Update(semesterId string, semester domains.Semester) error {
	err := repos.SemesterRepo.Update(semesterId, semester)
	return err
}

func (ss *semesterService) Publish(semesterIds []string) error {
	err := repos.SemesterRepo.Publish(semesterIds)
	return err
}

func (ss *semesterService) Delete(semesterId string) error {
	err := repos.SemesterRepo.Delete(semesterId)
	return err
}
