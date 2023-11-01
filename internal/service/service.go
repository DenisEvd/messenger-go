package service

import (
	"messenger-go/domain"
	"messenger-go/internal/repository"
)

type Message interface {
	Create(message domain.Message) (int, error)
	GetAll(senderID int, receiverID int) ([]domain.Message, error)
}

type Service struct {
	Message
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Message: repo.Message,
	}
}
