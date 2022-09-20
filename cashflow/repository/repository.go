package repository

import (
	"github.com/Vladosya/our_project/appl_row"
	"github.com/jmoiron/sqlx"
)

type TodoUser interface {
	RegistrationUser() error
}

type TodoAd interface {
	AdChangeParams(city string, price int) error
	CreateAd(adParam appl_row.Ad) error
	ActivateAd(id int) error
	ToCompleteAd(id int) error
	SummarizingAd() error
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
