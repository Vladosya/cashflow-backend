package repository

import (
	"github.com/jmoiron/sqlx"
)

type AdPostgres struct {
	db *sqlx.DB
}

func NewAdPostgres(db *sqlx.DB) *AdPostgres {
	return &AdPostgres{
		db: db,
	}
}

func (r *AdPostgres) AdChangeParams(city string, price int) error {
	_, err := r.db.Exec("UPDATE ad_params SET price = $1 WHERE city = $2", price, city)
	if err != nil {
		return err
	}
	return nil
}
