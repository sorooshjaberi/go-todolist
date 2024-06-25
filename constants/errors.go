package constants

import (
	"errors"
	"log"
)

var (
	ErrEmptyCredentials   = errors.New("username or password is empty")
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrUserNotFound       = errors.New("no such user exists")
	ErrInternalServer     = errors.New("internal server error")
	ErrUserAlreadyExists  = errors.New("user already exists")
)

func HandleErrorSoft(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func HandleErrorByPanic(err error) {
	if err != nil {
		panic(err)
	}
}
