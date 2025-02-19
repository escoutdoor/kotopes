package model

import "time"

type Favorite struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	PetID     string    `db:"pet_id"`
	CreatedAt time.Time `db:"created_at"`
}
