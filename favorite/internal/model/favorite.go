package model

import (
	"time"
)

type Favorite struct {
	ID        string
	UserID    string
	PetID     string
	CreatedAt time.Time
}

type FavoritePet struct {
	ID        string
	UserID    string
	Pet       *Pet
	CreatedAt time.Time
}

type Pet struct {
	ID          string
	Name        string
	Description string
	Age         int32
	OwnerID     string
	CreatedAt   time.Time
}

type CreateFavorite struct {
	PetID  string
	UserID string
}

type DeleteFavorite struct {
	FavoriteID string
	UserID     string
}

type ListFavorites struct {
	UserID string
	Limit  int32
	Offset int32
}
