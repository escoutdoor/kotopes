package model

import "time"

type Pet struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Age         int32     `db:"age"`
	OwnerID     string    `db:"owner_id"`
	CreatedAt   time.Time `db:"created_at"`
}
