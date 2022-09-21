package repository

import (
	"github.com/Vladosya/our_project/appl_row"
	"github.com/jmoiron/sqlx"
)

type TodoUser interface {
	RegistrationUser() (error, int)
	EntryToAd(userId int, adId int) (error, int)
	RefusalAd(userId int, adId int) (error, int)
}

type TodoAd interface {
	AdChangeParams(city string, price int) (error, int)
	CreateAd(adParam appl_row.Ad) (error, int)
	ActivateAd(id int) (error, int)
	ToCompleteAd(id int) (error, int)
	CancelAd(id int) (error, int)
	SummarizingAd(adId int, winnersPart []appl_row.WinnersPart) (error, int)
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
