package service

import (
	"github.com/Vladosya/our_project/appl_row"
	"github.com/Vladosya/our_project/cashflow/repository"
)

type TodoUser interface {
	RegistrationUser() (error, int)
	EntryToAd(userId int, adId int) (error, int)
}

type TodoAd interface {
	AdChangeParams(city string, price int) (error, int)
	CreateAd(adParam appl_row.Ad) (error, int)
	ActivateAd(id int) (error, int)
	ToCompleteAd(id int) (error, int)
	CancelAd(id int) (error, int)
	SummarizingAd() (error, int)
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
