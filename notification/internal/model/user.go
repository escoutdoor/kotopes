package model

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	City      string
	Country   string
	CreatedAt time.Time
}
