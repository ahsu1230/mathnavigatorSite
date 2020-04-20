package domains

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct {
	HashBytes []byte
}

func (hash *Hash) GetHash(str string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hash.HashBytes = bytes
	return nil
}

func (hash *Hash) Compare(str string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash.HashBytes, []byte(str))
	return err == nil, err
}
