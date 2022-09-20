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

func (s *AdService) AdChangeParams(city string, price int) error {
	return s.repo.AdChangeParams(city, price)
}

func (s *AdService) CreateAd(adParam appl_row.Ad) error {
	return s.repo.CreateAd(adParam)
}

func (s *AdService) ActivateAd(id int) error {
	return s.repo.ActivateAd(id)
}

func (s *AdService) ToCompleteAd(id int) error {
	return s.repo.ToCompleteAd(id)
}

func (s *AdService) CancelAd(id int) error {
	return s.repo.CancelAd(id)
}

func (s *AdService) SummarizingAd() error {
	return s.repo.SummarizingAd()
}
