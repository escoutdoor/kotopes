package model

import "time"

type FavoriteList struct {
	Favorites []*FavoritePet
	Total     int32
}

type FavoritePet struct {
	ID        string
	UserID    string
	Pet       *Pet
	CreatedAt time.Time
}

type Favorite struct {
	ID        string
	UserID    string
	PetID     string
	CreatedAt time.Time
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
