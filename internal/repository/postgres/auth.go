package postgres

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"messenger-go/domain"
)

var ErrUserAlreadyExists = errors.New("username already exists")

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthorizationPostgres(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{db: db}
}

func (a *AuthorizationPostgres) CreateUser(user domain.User) (int, error) {
	exists, err := a.isExistsUser(user.Username)
	if err != nil {
		return 0, errors.Wrap(err, "error creating user")
	}

	if exists {
		return 0, ErrUserAlreadyExists
	}

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

func (a *AuthorizationPostgres) isExistsUser(username string) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 LIMIT 1", usersTable)

	var tmp int
	err := a.db.QueryRow(query, username).Scan(&tmp)
	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, errors.Wrap(err, "error check is user exists in db")
	}

	return true, nil
}
