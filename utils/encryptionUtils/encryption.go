package encryptionUtils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
	"todolist/constants"
	"todolist/lib/dotenvLib"
	"todolist/utils/errorsUtils"
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

func GenerateJWT(username string, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    username,
		"userId": userId,
		"exp":    time.Now().Add(ExpirationDuration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(dotenvLib.GetEnv(constants.EnvKeys.JWTSecret).(string)))

	if err != nil {
		errorsUtils.HandleErrorSoft(err)
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", constants.ErrUnexpectedTokenMethod
		}
		return []byte(dotenvLib.GetEnv(constants.EnvKeys.JWTSecret).(string)), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", constants.ErrInvalidJWTToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", constants.ErrInvalidTokenClaims
	}

	fmt.Println(claims)

	if claims["userId"] == nil {
		return "", constants.ErrInvalidTokenClaims
	}

	userId := fmt.Sprint(claims["userId"])

	return userId, nil
}
