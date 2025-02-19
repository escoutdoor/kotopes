package model

type AccessCheck struct {
	Endpoint string
	Method   string
	UserID   string
	Role     string
}

type AccessInfo struct {
	IsAllowed bool
}
