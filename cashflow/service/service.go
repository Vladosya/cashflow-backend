package service

import (
	"github.com/Vladosya/our_project/appl_row"
	"github.com/Vladosya/our_project/cashflow/repository"
)

type TodoUser interface {
	RegistrationUser() error
	EntryToAd(userId int, adId int) (error, int)
}

type TodoAd interface {
	AdChangeParams(city string, price int) error
	CreateAd(adParam appl_row.Ad) error
	ActivateAd(id int) error
	ToCompleteAd(id int) error
	CancelAd(id int) error
	SummarizingAd() error
}

type Service struct {
	TodoUser
	TodoAd
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TodoUser: NewUserService(r.TodoUser),
		TodoAd:   NewAdService(r.TodoAd),
	}
}
