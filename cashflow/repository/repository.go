package repository

import (
	"github.com/jmoiron/sqlx"
)

type TodoUser interface {
	RegistrationUser() error
}

type TodoAd interface {
	AdChangeParams(city string, price int) error
}

type Repository struct {
	TodoUser
	TodoAd
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser: NewUserPostgres(db),
		TodoAd:   NewAdPostgres(db),
	}
}
