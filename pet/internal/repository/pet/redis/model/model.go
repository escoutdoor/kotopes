package model

import (
	"encoding/json"
	"time"
)

type Pet struct {
	ID          string    `redis:"id"`
	Name        string    `redis:"name"`
	Description string    `redis:"description"`
	Age         int32     `redis:"age"`
	OwnerID     string    `redis:"owner_id"`
	CreatedAt   time.Time `redis:"created_at"`
}

func (p *Pet) MarshalBinary() (data []byte, err error) {
	return json.Marshal(p)
}

func (p *Pet) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}
