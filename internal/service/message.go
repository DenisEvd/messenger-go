package service

import (
	"messenger-go/domain"
	"messenger-go/internal/repository"
)

type MessageService struct {
	repository repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repository: repo}
}

func (m *MessageService) Create(message domain.Message) (int, error) {
	return 0, nil
}

func (m *MessageService) GetAll(senderID int, receiverID int) ([]domain.Message, error) {
	return []domain.Message{}, nil
}
