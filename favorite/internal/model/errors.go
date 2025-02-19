package model

import "errors"

var (
	ErrFavoriteNotFound   = errors.New("favorite not found")
	ErrAlreadyInList      = errors.New("already in the favorite list")
	ErrNotOwnerOfFavorite = errors.New("not the owner of this favorite item")
)
