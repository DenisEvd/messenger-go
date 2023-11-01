package domain

import "time"

type Message struct {
	SenderID   int    `json:"sender_id" binding:"required"`
	ReceiverID int    `json:"receiver_id" binding:"required"`
	Text       string `json:"text" binding:"required"`
	Timestamp  time.Time
}
