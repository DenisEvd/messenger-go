package service

import (
	"github.com/pkg/errors"
	"messenger-go/domain"
	"messenger-go/internal/repository"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo: repo}
}

func (c *ChatService) Create(name string) (int, error) {
	id, err := c.repo.Create(name)
	if err != nil {
		return 0, errors.Wrap(err, "error chat service")
	}

	return id, nil
}

func (c *ChatService) GetUserChats(userID int) ([]domain.Chat, error) {
	chats, err := c.repo.GetUserChats(userID)
	if err != nil {
		return []domain.Chat{}, errors.Wrap(err, "error chat service")
	}

	return chats, nil
}

func (c *ChatService) AddUser(chatID int, userID int) error {
	if err := c.repo.AddUser(chatID, userID); err != nil {
		return errors.Wrap(err, "error chat service")
	}

	return nil
}
