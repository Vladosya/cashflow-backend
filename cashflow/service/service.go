package service

import (
	"github.com/Vladosya/our_project/cashflow/repository"
)

type TodoUser interface {
	RegistrationUser() error
}

type Service struct {
	TodoUser
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TodoUser: NewUserService(r.TodoUser),
	}
}
