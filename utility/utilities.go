package utility

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "Failed to convert!", err
	}
	return string(hashPass), nil
}

func ComparePassward(password string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
