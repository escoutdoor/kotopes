package model

import "errors"

var (
	ErrUserAlreadyExists      = errors.New("user already exists")
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)
