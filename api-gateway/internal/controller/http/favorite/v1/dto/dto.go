package dto

import (
	"time"
)

type ListFavoritesResponse struct {
	Favorites []*FavoritePet `json:"favorites"`
	Total     int32          `json:"total"`
}

type FavoritePet struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Pet       *Pet      `json:"pet"`
	CreatedAt time.Time `json:"created_at"`
}

type Pet struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Age         int32     `json:"age"`
	OwnerID     string    `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
}
