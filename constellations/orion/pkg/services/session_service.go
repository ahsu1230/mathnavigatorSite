package services

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/repos"
)

var SessionService sessionServiceInterface = &sessionService{}

// Interface for SessionService
type sessionServiceInterface interface {
	GetAllByClassId(string, bool) ([]domains.Session, error)
	GetAllUnpublished() ([]domains.Session, error)
	GetBySessionId(uint) (domains.Session, error)
	Create([]domains.Session) error
	Update(uint, domains.Session) error
	Publish([]uint) error
	Delete([]uint) error
}

// Struct that implements interface
type sessionService struct{}

func (ss *sessionService) GetAllByClassId(classId string, publishedOnly bool) ([]domains.Session, error) {
	sessions, err := repos.SessionRepo.SelectAllByClassId(classId, publishedOnly)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (ss *sessionService) GetAllUnpublished() ([]domains.Session, error) {
	sessions, err := repos.SessionRepo.SelectAllUnpublished()
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (ss *sessionService) GetBySessionId(id uint) (domains.Session, error) {
	session, err := repos.SessionRepo.SelectBySessionId(id)
	if err != nil {
		return domains.Session{}, err
	}
	return session, nil
}

func (ss *sessionService) Create(sessions []domains.Session) error {
	err := repos.SessionRepo.Insert(sessions)
	return err
}

func (ss *sessionService) Update(id uint, session domains.Session) error {
	err := repos.SessionRepo.Update(id, session)
	return err
}

func (ss *sessionService) Publish(ids []uint) error {
	err := repos.SessionRepo.Publish(ids)
	return err
}

func (ss *sessionService) Delete(ids []uint) error {
	err := repos.SessionRepo.Delete(ids)
	return err
}
