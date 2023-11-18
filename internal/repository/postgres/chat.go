package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"messenger-go/domain"
)

var ErrNoChatFound = errors.New("data base doesn't exist chat")

type ChatPostgres struct {
	db *sqlx.DB
}

func NewChatPostgres(db *sqlx.DB) *ChatPostgres {
	return &ChatPostgres{db: db}
}

func (c *ChatPostgres) Create(name string) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", chatsTable)

	var id int
	if err := c.db.QueryRow(query, name).Scan(&id); err != nil {
		return 0, errors.Wrap(err, "error inserting chat")
	}

	return id, nil
}

func (c *ChatPostgres) GetUserChats(userID int) ([]domain.Chat, error) {
	query := fmt.Sprintf("SELECT c.id, c.name FROM %s c INNER JOIN %s cu ON c.id = cu.chat_id WHERE cu.user_id = $1", chatsTable, chatsUsersTable)

	rows, err := c.db.Query(query, userID)
	if err != nil {
		return []domain.Chat{}, errors.Wrap(err, "error selecting chats")
	}
	defer func() { _ = rows.Close() }()

	chats := make([]domain.Chat, 0)
	for rows.Next() {
		var chat domain.Chat
		if err := rows.Scan(&chat.ID, &chat.Name); err != nil {
			return []domain.Chat{}, errors.Wrap(err, "error getting chats")
		}
		chats = append(chats, chat)
	}

	return chats, nil
}

func (c *ChatPostgres) AddUser(chatID int, userID int) error {
	existsChat, err := c.isExistsChat(chatID)
	if err != nil {
		return err
	}

	if !existsChat {
		return ErrNoChatFound
	}

	query := fmt.Sprintf("INSERT INTO %s (chat_id, user_id) VALUES ($1, $2)", chatsUsersTable)

	if _, err := c.db.Exec(query, chatID, userID); err != nil {
		return err
	}

	return nil
}

func (c *ChatPostgres) isExistsChat(chatID int) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE id = $1", chatsTable)

	var count int
	if err := c.db.QueryRow(query, chatID).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}
