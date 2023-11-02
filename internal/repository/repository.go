package repository

import "messenger-go/domain"

type Message interface {
	Create(message domain.Message) (int, error)
	GetAll(userID int, chatID int) ([]domain.Message, error)
}

type Repository struct {
	Message
}

func NewRepository(message Message) *Repository {
	return &Repository{Message: message}
}
