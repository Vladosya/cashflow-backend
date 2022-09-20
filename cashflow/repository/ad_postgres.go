package repository

import (
	"fmt"
	"github.com/Vladosya/our_project/appl_row"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type AdPostgres struct {
	db *sqlx.DB
}

func NewAdPostgres(db *sqlx.DB) *AdPostgres {
	return &AdPostgres{
		db: db,
	}
}

func (r *AdPostgres) AdChangeParams(city string, price int) (error, int) {
	_, err := r.db.Exec("UPDATE ad_params SET price = $1 WHERE city = $2", price, city)
	if err != nil {
		return fmt.Errorf("ошибка изменения из базы данных, %s", err), http.StatusInternalServerError
	}
	return fmt.Errorf("успешное изменение стоимости участия в городе %s", city), http.StatusOK
}

func (r *AdPostgres) CreateAd(adParam appl_row.Ad) (error, int) {
	_, err := r.db.Exec("INSERT INTO ad (title, date_start, city, price, description, event_type, serial_number, points_options, is_visible) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		adParam.Title, adParam.DateStart, adParam.City,
		adParam.Price, adParam.Description, adParam.EventType,
		adParam.SerialNumber, adParam.PointsOptions, adParam.IsVisible,
	)
	if err != nil {
		return fmt.Errorf("ошибка создания из базы данных, %s", err), http.StatusInternalServerError
	}
	return fmt.Errorf("успешное создание мероприятия"), http.StatusOK
}

func (r *AdPostgres) ActivateAd(id int) (error, int) {
	_, err := r.db.Exec("UPDATE ad SET is_visible = true WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("ошибка обновления из базы данных, %s", err), http.StatusInternalServerError
	}
	return fmt.Errorf("успешная активация мероприятия"), http.StatusOK
}

func (r *AdPostgres) ToCompleteAd(id int) (error, int) {
	_, err := r.db.Exec("UPDATE ad SET is_finished = true WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("ошибка обновления из базы данных, %s", err), http.StatusInternalServerError
	}
	return fmt.Errorf("успешное завершение мероприятия"), http.StatusOK
}

func (r *AdPostgres) CancelAd(id int) (error, int) {
	_, err := r.db.Exec("UPDATE ad SET is_cancel = true WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("ошибка обновления из базы данных, %s", err), http.StatusInternalServerError
	}
	return fmt.Errorf("успешная отмена мероприятия"), http.StatusOK
}

func (r *AdPostgres) SummarizingAd() (error, int) {
	return fmt.Errorf("успешное распределение баллов"), http.StatusOK
}
