package repository

import (
	"github.com/jmoiron/sqlx"
)

type TodoUser interface {
	RegistrationUser() error
}

type Repository struct {
	TodoUser
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser: NewUserPostgres(db),
	}
}
