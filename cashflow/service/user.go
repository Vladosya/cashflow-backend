package service

import (
	"github.com/Vladosya/our_project/cashflow/repository"
)

type UserService struct {
	repo repository.TodoUser
}

func NewUserService(r repository.TodoUser) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) RegistrationUser() (error, int) {
	return s.repo.RegistrationUser()
}

func (s *UserService) EntryToAd(userId int, adId int) (error, int) {
	return s.repo.EntryToAd(userId, adId)
}

func (s *UserService) RefusalAd(userId int, adId int) (error, int) {
	return s.repo.RefusalAd(userId, adId)
}
