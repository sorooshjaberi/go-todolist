package utils

import "golang.org/x/crypto/bcrypt"

func HashText(text string) (string, error) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedText), nil
}

func VerifyHashedText(hashedText string, originalText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(originalText))
	return err == nil
}
