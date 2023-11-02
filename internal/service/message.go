package service

import (
	"github.com/pkg/errors"
	"messenger-go/domain"
	"messenger-go/internal/repository"
	"time"
)

type MessageService struct {
	repository repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repository: repo}
}

func (m *MessageService) Create(message domain.Message) (int, error) {
	message.Timestamp = time.Now()

	id, err := m.repository.Create(message)
	if err != nil {
		return 0, errors.Wrap(err, "error service message creating")
	}

	return id, nil
}

func (m *MessageService) GetAll(userID int, chatID int) ([]domain.Message, error) {
	messages, err := m.repository.GetAll(userID, chatID)
	if err != nil {
		return []domain.Message{}, errors.Wrap(err, "error getting all messages")
	}

	return messages, nil
}
