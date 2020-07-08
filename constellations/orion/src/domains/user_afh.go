package domains

import "errors"

var TABLE_USERAFH = "user_afh"

type UserAfh struct {
	Id     uint `json:"id"`
	UserId uint `json:"user_id"`
	AfhId  uint `json:"afh_id"`
}

func (userAfh *UserAfh) Validate() error {
	userId := userAfh.UserId
	afhId := userAfh.AfhId

	// UserId validation
	if userId < 0 {
		return errors.New("invalid userId")
	}

	// AfhId validation
	if afhId < 0 {
		return errors.New("invalid afhId")
	}

	return nil
}
