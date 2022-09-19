package service

import (
	"github.com/Vladosya/our_project/cashflow/repository"
)

type TodoUser interface {
	RegistrationUser() error
}

type TodoAd interface {
	AdChangeParams(city string, price int) error
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
