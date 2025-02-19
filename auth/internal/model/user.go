package model

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      string
	Phone     string
	City      string
	Country   string
	CreatedAt time.Time
}

type UpdateUser struct {
	ID        string
	FirstName *string
	LastName  *string
	Email     *string
	Password  *string
	Phone     *string
	City      *string
	Country   *string
}

type CreateUser struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}
