package utils

import "golang.org/x/crypto/bcrypt"

func BcryptHash(str string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	return string(bytes)
}

func BcryptCheck(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))

	return err == nil
}
