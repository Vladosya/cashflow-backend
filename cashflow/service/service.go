package service

import (
	"github.com/Vladosya/our_project/appl_row"
	"github.com/Vladosya/our_project/cashflow/repository"
)

type TodoSupport interface {
	EntryToAdSupport(userId int, adId int) (error, int)
}

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
	ReplantAd(adId int, seatAtTables []appl_row.SeatAtTables) (error, int)
	GetAllAd() ([]appl_row.AdFull, error, int)
}

type Service struct {
	TodoSupport
	TodoUser
	TodoAd
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TodoSupport: NewSupportService(r.TodoSupport),
		TodoUser:    NewUserService(r.TodoUser),
		TodoAd:      NewAdService(r.TodoAd),
	}
}
