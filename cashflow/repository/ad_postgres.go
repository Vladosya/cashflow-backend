package repository

import (
	"github.com/Vladosya/our_project/appl_row"
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

func (r *AdPostgres) CreateAd(adParam appl_row.Ad) error {
	_, err := r.db.Exec("INSERT INTO ad (title, date_start, city, price, description, event_type, serial_number, points_options, is_visible) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		adParam.Title, adParam.DateStart, adParam.City,
		adParam.Price, adParam.Description, adParam.EventType,
		adParam.SerialNumber, adParam.PointsOptions, adParam.IsVisible,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *AdPostgres) SummarizingAd() error {
	return nil
}
