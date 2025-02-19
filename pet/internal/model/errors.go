package model

import "errors"

var (
	ErrPetNotFound         = errors.New("pet not found")
	ErrNotPetOwner         = errors.New("you must be the owner of the pet to do this")
	ErrNoFieldsForUpdating = errors.New("there are no fields for updating")
)
