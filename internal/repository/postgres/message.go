package postgres

import "messenger-go/domain"

type MessagePostgres struct {
}

func NeMessagePostgres() *MessagePostgres {
	return &MessagePostgres{}
}

func (m *MessagePostgres) Create(message domain.Message) (int, error) {
	return 0, nil
}

func (m *MessagePostgres) GetAll(senderID int, receiverID int) ([]domain.Message, error) {
	return []domain.Message{}, nil
}
