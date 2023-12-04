package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	auth "github.com/eeQuillibrium/go-rest-auth"
)

type authPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *authPostgres {
	return &authPostgres{db: db}
}

func (a *authPostgres) CreateUser(user auth.User) (int, error) {
	var id int
	q := fmt.Sprintf("INSERT INTO %s (login, password) values($1, $2) returning id", users)
	rows := a.db.QueryRow(q, user.Login, user.Password)
	if err := rows.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *authPostgres) CheckUser(user auth.User) (int, error) {

	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2", users)

	row := a.db.QueryRow(query, user.Login, user.Password)

	var id string

	if err := row.Scan(&id); err != nil {
		return -2, err
	}

	var err error
	user.Id, err = strconv.Atoi(id)
	if err != nil {
		return -1, nil
	}
	
	return user.Id, err
}
