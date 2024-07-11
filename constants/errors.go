package constants

import (
	"errors"
)

var (
	ErrEmptyCredentials      = errors.New("username or password is empty")
	ErrInvalidCredentials    = errors.New("invalid username or password")
	ErrUserNotFound          = errors.New("no such user exists")
	ErrTodoNotFound          = errors.New("no such todo exists")
	ErrInternalServer        = errors.New("internal server error")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrInvalidJWTToken       = errors.New("invalid jwt token")
	ErrMissingToken          = errors.New("missing token")
	ErrUnexpectedTokenMethod = errors.New("unexpected token method")
	ErrInvalidTokenClaims    = errors.New("invalid token claims")
	ErrInvalidPayload        = errors.New("invalid payload")
	ErrIdShouldBeNumber      = errors.New("id should be a number")
)
