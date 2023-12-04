package repository

import (
	"database/sql"

	auth "github.com/eeQuillibrium/go-rest-auth"
	_ "github.com/lib/pq"
)

type Authorization interface {
	CreateUser(user auth.User) (int, error)
	CheckUser(user auth.User) (int, error)
}
type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Authorization: NewAuthPostgres(db)}
}
