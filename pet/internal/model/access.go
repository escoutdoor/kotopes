package model

type CheckAccess struct {
	Endpoint string
	UserID   string
	Role     string
}

type AccessInfo struct {
	IsAllowed bool
}
