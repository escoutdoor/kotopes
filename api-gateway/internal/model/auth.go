package model

type Login struct {
	Email    string
	Password string
}

type Register struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     *string
}

type Token struct {
	UserID string
	Role   string
}

type AuthTokens struct {
	AccessToken  string
	RefreshToken string
}
