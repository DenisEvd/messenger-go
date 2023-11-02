package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	messagesTable = "messages"
	chatsTable    = "chats"
	usersTable    = "users"
)

func NewPostgres(conf *Config) (*sqlx.DB, error) {
	fmt.Println(conf)
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.Host, conf.Port, conf.Username, conf.Password, conf.DBName, conf.SSLMode))
	if err != nil {
		return nil, errors.Wrap(err, "error postgres connect")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "error postgres connection")
	}

	return db, nil
}
