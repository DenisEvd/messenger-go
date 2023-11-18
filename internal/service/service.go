package service

import (
	"messenger-go/domain"
	"messenger-go/internal/repository"
)

type Message interface {
	Create(message domain.Message) (int, error)
	GetAll(userID int, chatID int) ([]domain.Message, error)
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Chat interface {
	Create(name string) (int, error)
	GetUserChats(userID int) ([]domain.Chat, error)
	AddUser(chatID int, userID int) error
}

type Service struct {
	Message
	Authorization
	Chat
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Message:       NewMessageService(repo.Message),
		Authorization: NewAuthorizationService(repo.Authorization),
		Chat:          NewChatService(repo.Chat),
	}
}
