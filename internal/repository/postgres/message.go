package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"messenger-go/domain"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessagePostgres(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (m *MessagePostgres) Create(message domain.Message) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, chat_id, message, sending_time) VALUES($1, $2, $3, $4) RETURNING id", messagesTable)
	if err := m.db.QueryRow(query, message.UserID, message.ChatID, message.Text, message.Timestamp).Scan(&id); err != nil {
		return 0, errors.Wrap(err, "error inserting data")
	}

	return id, nil
}

func (m *MessagePostgres) GetAll(userID int, chatID int) ([]domain.Message, error) {
	query := fmt.Sprintf("SELECT user_id, chat_id, message, sending_time FROM %s WHERE user_id=$1 AND chat_id=$2", messagesTable)
	rows, err := m.db.Query(query, userID, chatID)
	if err != nil {
		return []domain.Message{}, errors.Wrap(err, "error selecting all messages")
	}
	defer func() { _ = rows.Close() }()

	messages := make([]domain.Message, 0)
	for rows.Next() {
		var message domain.Message
		if err := rows.Scan(&message.UserID, &message.ChatID, &message.Text, &message.Timestamp); err != nil {
			return []domain.Message{}, errors.Wrap(err, "error scan all messages")
		}
		messages = append(messages, message)
	}

	return messages, nil
}
