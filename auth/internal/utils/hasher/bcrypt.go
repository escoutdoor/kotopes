package hasher

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct {
}

func NewBcryptHasher() Hasher {
	return &bcryptHasher{}
}

func (h *bcryptHasher) Compare(pw, hashPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPw), []byte(pw))
	if err != nil {
		return false
	}

	return true
}

func (h *bcryptHasher) Hash(pw string) (string, error) {
	hashPw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("generate hash from password error: %s", err)
	}

	return string(hashPw), nil
}
