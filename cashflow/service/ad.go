package service

import (
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
