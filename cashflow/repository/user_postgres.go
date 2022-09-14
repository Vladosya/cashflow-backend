package repository

import (
	"fmt"
	"github.com/Vladosya/our_project/helpers"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) RegistrationUser() error {
	fmt.Println("RegistrationUser")
	fmt.Println("helpers.RandomStrGeneration(8)", helpers.RandomStrGeneration(8))
	return nil
}
