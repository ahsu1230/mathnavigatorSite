package domains

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct {
	HashBytes []byte
}

func NewHash(str string) (Hash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return Hash{}, err
	}
	return Hash{bytes}, nil
}

func (hash *Hash) Compare(str string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash.HashBytes, []byte(str))
	return err == nil, err
}
