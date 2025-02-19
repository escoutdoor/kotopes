package model

import "time"

type ChatMessageNotification struct {
	ChatID  string
	Message Message
}

type Message struct {
	ID        string
	UserID    string
	Content   string
	CreatedAt time.Time
}
