package encryptionUtil

import (
	"booking/constants"
	"booking/lib/dotenvLib"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

var (
	ExpirationDuration = time.Hour * 24
)

func HashText(text string) (string, error) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hashedText), nil
}

func VerifyHashedText(hashedText string, originalText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(originalText))
	return err == nil
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(ExpirationDuration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(dotenvLib.GetEnv(constants.EnvKeys.JWTSecret).(string)))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(dotenvLib.GetEnv(constants.EnvKeys.JWTSecret).(string)), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", constants.ErrInvalidJWTToken
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}
	return sub, nil
}
