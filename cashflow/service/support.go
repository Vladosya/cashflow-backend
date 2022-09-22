package service

import "github.com/Vladosya/our_project/cashflow/repository"

type SupportService struct {
	repo repository.TodoSupport
}

func NewSupportService(r repository.TodoSupport) *SupportService {
	return &SupportService{
		repo: r,
	}
}

func (s *SupportService) EntryToAdSupport(userId int, adId int) (error, int) {
	return s.repo.EntryToAdSupport(userId, adId)
}
