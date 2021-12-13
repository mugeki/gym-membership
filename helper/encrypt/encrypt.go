package encrypt

import "golang.org/x/crypto/bcrypt"

type Helper interface{
	Hash(password string) (string, error)
	ValidateHash(password, hash string) bool
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func ValidateHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}