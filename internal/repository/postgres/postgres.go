package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
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

func NewPostgres(conf Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.Host, conf.Port, conf.Username, conf.Password, conf.DBName, conf.SSLMode))
	if err != nil {
		return nil, errors.Wrap(err, "error postgres connect")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "error postgres connection")
	}

	return db, nil
}
