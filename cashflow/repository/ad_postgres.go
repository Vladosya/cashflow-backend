package repository

import (
	"encoding/json"
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

func (r *AdPostgres) SummarizingAd(adId int, winnersPart []appl_row.WinnersPart) (error, int) {
	rowAd, err := r.db.Query("SELECT * FROM ad WHERE id = $1", adId)
	if err != nil {
		return fmt.Errorf("ошибка получения из базы данных, %s", err), http.StatusInternalServerError
	}
	defer rowAd.Close()
	var ad []Ad
	for rowAd.Next() {
		var p Ad
		if err := rowAd.Scan(
			&p.Id, &p.Title, &p.DateStart, &p.Created, &p.City,
			&p.Price, &p.Description, &p.EventType, &p.Participant, &p.SerialNumber,
			&p.PointOptions, &p.IsVisible, &p.IsFinished, &p.IsCancel,
		); err != nil {
			return fmt.Errorf("ошибка преобразования полученных данных, %s", err), http.StatusInternalServerError
		}
		ad = append(ad, p)
	}
	if len(ad) == 0 {
		return fmt.Errorf("данного мероприятия не существует"), http.StatusBadRequest
	} else {
		rowPointsGame, err := r.db.Query("SELECT id, title, city, version, scoring FROM points_game WHERE id = $1", ad[0].PointOptions)
		if err != nil {
			return fmt.Errorf("ошибка получения из базы данных, %s", err), http.StatusInternalServerError
		}
		defer rowPointsGame.Close()
		var pointsGame []appl_row.PointsGame
		for rowPointsGame.Next() {
			var p appl_row.PointsGame
			if err := rowPointsGame.Scan(
				&p.Id, &p.Title, &p.City, &p.Version, &p.Scoring,
			); err != nil {
				return fmt.Errorf("ошибка преобразования полученных данных, %s", err), http.StatusInternalServerError
			}
			pointsGame = append(pointsGame, p)
		}
		var scoring appl_row.Scoring
		dataScoring := pointsGame[0].Scoring
		err = json.Unmarshal(dataScoring, &scoring)
		if err != nil {
			return fmt.Errorf("ошибка работы с json, %s", err), http.StatusInternalServerError
		}
		for i := 0; i < len(winnersPart); i++ { // записываем каждому игроку заработанное кол-во очков
			for j := 0; j < len(winnersPart[i].WinUser); j++ {
				winnersPart[i].WinUser[j].Assigned = scoring.WinRes[j].NumberPoints
			}
		}
		fmt.Println("winnersPart -->", winnersPart)
		// ОСТАЛОСЬ пробежаться по пользователям и добавить им
	}
	return fmt.Errorf("успешное распределение баллов"), http.StatusOK
}
