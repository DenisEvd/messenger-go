package postgres

import (
	"github.com/jmoiron/sqlx"
	"messenger-go/domain"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NeMessagePostgres(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (m *MessagePostgres) Create(message domain.Message) (int, error) {
	return 0, nil
}

func (m *MessagePostgres) GetAll(senderID int, receiverID int) ([]domain.Message, error) {
	return []domain.Message{}, nil
}
