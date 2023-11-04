package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"messenger-go/domain"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthorizationPostgres(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{db: db}
}

func (a *AuthorizationPostgres) CreateUser(user domain.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password) VALUES ($1, $2, $3) RETURNING id", usersTable)

	var id int
	if err := a.db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&id); err != nil {
		return 0, errors.Wrap(err, "error creating user")
	}

	return id, nil
}

func (a *AuthorizationPostgres) GetUser(username string, password string) (domain.User, error) {
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE username=$1 AND password=$2", usersTable)

	user := domain.User{
		Username: username,
		Password: password,
	}

	if err := a.db.QueryRow(query, username, password).Scan(&user.ID, &user.Name); err != nil {
		return domain.User{}, errors.Wrap(err, "error getting user")
	}

	return user, nil
}
