package dto

import "time"

type CreatePetRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Age         int32  `json:"age"`
}

type CreatePetResponse struct {
	ID string `json:"id"`
}

type UpdatePetRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Age         *int32  `json:"age,omitempty"`
}

type GetPetResponse struct {
	Pet *Pet `json:"pet"`
}

type ListPetsResponse struct {
	Pets []*Pet `json:"pets"`
}

type Pet struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Age         int32     `json:"age"`
	OwnerID     string    `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
}
