package model

type FavoriteMessage struct {
	OwnerID string `json:"owner_id"`
	PetID   string `json:"pet_id"`
	UserID  string `json:"user_id"`
}
