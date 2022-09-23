package service

import (
	"github.com/Vladosya/our_project/appl_row"
	"github.com/Vladosya/our_project/cashflow/repository"
)

type AdService struct {
	repo repository.TodoAd
}

func NewAdService(r repository.TodoAd) *AdService {
	return &AdService{
		repo: r,
	}
}

func (s *AdService) AdChangeParams(city string, price int) (error, int) {
	return s.repo.AdChangeParams(city, price)
}

func (s *AdService) CreateAd(adParam appl_row.Ad) (error, int) {
	return s.repo.CreateAd(adParam)
}

func (s *AdService) ActivateAd(id int) (error, int) {
	return s.repo.ActivateAd(id)
}

func (s *AdService) ToCompleteAd(id int) (error, int) {
	return s.repo.ToCompleteAd(id)
}

func (s *AdService) CancelAd(id int) (error, int) {
	return s.repo.CancelAd(id)
}

func (s *AdService) SummarizingAd(adId int, winnersPart []appl_row.WinnersPart) (error, int) {
	return s.repo.SummarizingAd(adId, winnersPart)
}

func (s *AdService) ReplantAd(adId int, seatAtTables []appl_row.SeatAtTables) (error, int) {
	return s.repo.ReplantAd(adId, seatAtTables)
}

func (s *AdService) GetInfoAbTables(adId int) ([]appl_row.GameForm, error, int) {
	return s.repo.GetInfoAbTables(adId)
}

func (s *AdService) GetAllAd() ([]appl_row.AdFull, error, int) {
	return s.repo.GetAllAd()
}
