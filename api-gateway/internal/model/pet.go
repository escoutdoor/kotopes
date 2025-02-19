package model

import "time"

type Pet struct {
	ID          string
	Name        string
	Description string
	Age         int32
	OwnerID     string
	CreatedAt   time.Time
}

type CreatePet struct {
	Name        string
	Description string
	Age         int32
	OwnerID     string
}

type UpdatePet struct {
	ID          string
	OwnerID     string
	Name        *string
	Description *string
	Age         *int32
}

type ListPets struct {
	Limit  int32
	Offset int32
	PetIDs []string
}
