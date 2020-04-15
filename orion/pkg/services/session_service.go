package services

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
)

var SessionService sessionServiceInterface = &sessionService{}

// Interface for SessionService
type sessionServiceInterface interface {
	GetAllByClassId(string, bool) ([]domains.Session, error)
	GetAllUnpublished() ([]domains.Session, error)
	GetBySessionId(uint) (domains.Session, error)
	Create(domains.Session) error
	Publish([]uint) []domains.PublishErrorBody
	Update(uint, domains.Session) error
	Delete(uint) error
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

func (ss *sessionService) Create(session domains.Session) error {
	err := repos.SessionRepo.Insert(session)
	return err
}

func (ss *sessionService) Publish(ids []uint) []domains.PublishErrorBody {
	errors := repos.SessionRepo.Publish(ids)
	return errors
}

func (ss *sessionService) Update(id uint, session domains.Session) error {
	err := repos.SessionRepo.Update(id, session)
	return err
}

func (ss *sessionService) Delete(id uint) error {
	err := repos.SessionRepo.Delete(id)
	return err
}
