package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string         `db:"id"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	Role      string         `db:"role"`
	Phone     sql.NullString `db:"phone"`
	City      sql.NullString `db:"city"`
	Country   sql.NullString `db:"country"`
	CreatedAt time.Time      `db:"created_at"`
}
