package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type Login struct {
	Email    string
	Password string
}

type UserClaims struct {
	jwt.RegisteredClaims
	ID   string
	Role string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
