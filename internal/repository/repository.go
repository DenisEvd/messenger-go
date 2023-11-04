package repository

import "messenger-go/domain"

type Message interface {
	Create(message domain.Message) (int, error)
	GetAll(userID int, chatID int) ([]domain.Message, error)
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username string, password string) (domain.User, error)
}

type Repository struct {
	Message
	Authorization
}

func NewRepository(message Message, authorization Authorization) *Repository {
	return &Repository{
		Message:       message,
		Authorization: authorization,
	}
}
