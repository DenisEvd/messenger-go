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

type Chat interface {
	Create(name string) (int, error)
	GetUserChats(userID int) ([]domain.Chat, error)
	AddUser(chatID int, userID int) error
}

type Repository struct {
	Message
	Authorization
	Chat
}

func NewRepository(message Message, authorization Authorization, chat Chat) *Repository {
	return &Repository{
		Message:       message,
		Authorization: authorization,
		Chat:          chat,
	}
}
