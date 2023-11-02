package domain

import "time"

type Message struct {
	UserID    int    `json:"user_id" binding:"required"`
	ChatID    int    `json:"chat_id" binding:"required"`
	Text      string `json:"text" binding:"required"`
	Timestamp time.Time
}
