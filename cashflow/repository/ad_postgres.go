package repository

import (
	"encoding/json"
	"fmt"
	"github.com/Vladosya/our_project/appl_row"
	"github.com/Vladosya/our_project/helpers"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

type AdPostgres struct {
	db *sqlx.DB
}

func NewAdPostgres(db *sqlx.DB) *AdPostgres {
	return &AdPostgres{
		db: db,
	}
}

func (r *AdPostgres) AdChangeParams(city string, price int) (error, int) { // Изменение стоимости участия мероприятия
	_, err := r.db.Exec("UPDATE ad_params SET price = $1 WHERE city = $2", price, city)
	if err != nil {
		return fmt.Errorf("ошибка изменения из базы данных, %s", err), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (r *AdPostgres) CreateAd(adParam appl_row.Ad) (error, int) { // Создание мероприятия
	weekday := time.Now().Weekday()
	if int(weekday) == 3 || int(weekday) == 7 {
		var resGenSeatAtTableByTableLen = helpers.GenSeatAtTableByTableLen(adParam.LimitationTables)
		var adCreatedId = 0
		err := r.db.QueryRow("INSERT INTO ad (title, date_start, city, price, description, event_type, serial_number, points_options, is_visible, limitation_tables) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
			adParam.Title, adParam.DateStart, adParam.City,
			adParam.Price, adParam.Description, adParam.EventType,
			adParam.SerialNumber, adParam.PointsOptions, adParam.IsVisible,
			adParam.LimitationTables,
		).Scan(&adCreatedId)
		if err != nil {
			return fmt.Errorf("ошибка создания из базы данных и получение в Scan, %s", err), http.StatusInternalServerError
		}
		_, err = r.db.Exec("INSERT INTO game (id_ad, seat_at_table) VALUES($1, $2)", adCreatedId, resGenSeatAtTableByTableLen)
		if err != nil {
			return fmt.Errorf("ошибка создания из базы данных, %s", err), http.StatusInternalServerError
		}
	} else {
		return fmt.Errorf("создание мероприятий разрешено по средам и воскресеньям"), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (r *AdPostgres) ActivateAd(id int) (error, int) { // Активация мероприятия (сделать видимым мероприятие для всех пользователей)
	_, err := r.db.Exec("UPDATE ad SET is_visible = true WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("ошибка обновления из базы данных, %s", err), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (r *AdPostgres) ToCompleteAd(id int) (error, int) { // Завершить мероприятия
	_, err := r.db.Exec("UPDATE ad SET is_finished = true WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("ошибка обновления из базы данных, %s", err), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (r *AdPostgres) CancelAd(id int) (error, int) { // Отмена мероприятия (если мероприятие по каким-то причинам было отменено)
	_, err := r.db.Exec("UPDATE ad SET is_cancel = true WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("ошибка обновления из базы данных, %s", err), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (r *AdPostgres) SummarizingAd(adId int, winnersPart []appl_row.WinnersPart) (error, int) { // Распределение баллов по участникам за пройденное мероприятие
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
			&p.PointOptions, &p.IsVisible, &p.IsFinished, &p.IsCancel, &p.LimitationTables,
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
	return nil, http.StatusOK
}

func (r *AdPostgres) ReplantAd(adId int, seatAtTables []appl_row.SeatAtTables) (error, int) { // Пересаживание пользователей из одного стола в другой
	jsonData, err := json.Marshal(seatAtTables)
	if err != nil {
		return fmt.Errorf("ошибка при кодировки данных в JSON, %s", err), http.StatusInternalServerError
	}
	_, err = r.db.Exec("UPDATE game SET seat_at_table = $1 WHERE id_ad = $2", jsonData, adId)
	if err != nil {
		return fmt.Errorf("ошибка обновления из базы данных, %s", err), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (r *AdPostgres) GetInfoAbTables(adId int) ([]appl_row.GameForm, error, int) { // Получить данные о том, кто за каким столом сидит и т.д
	rowGame, err := r.db.Query("SELECT * FROM game WHERE id_ad = $1", adId)
	if err != nil {
		return []appl_row.GameForm{}, fmt.Errorf("ошибка получения из базы данных, %s", err), http.StatusInternalServerError
	}
	defer rowGame.Close()
	var game []appl_row.Game
	for rowGame.Next() {
		var p appl_row.Game
		if err := rowGame.Scan(
			&p.Id, &p.IdAd, &p.SeatAtTable, &p.IsFinished,
		); err != nil {
			return []appl_row.GameForm{}, fmt.Errorf("ошибка преобразования полученных данных, %s", err), http.StatusInternalServerError
		}
		game = append(game, p)
	}
	var formGame []appl_row.GameForm
	var formSeatAtTable []appl_row.SeatAtTables
	json.Unmarshal([]byte(game[0].SeatAtTable), &formSeatAtTable)
	formGame = append(formGame, appl_row.GameForm{
		Id:          game[0].Id,
		IdAd:        game[0].IdAd,
		SeatAtTable: formSeatAtTable,
		IsFinished:  game[0].IsFinished,
	})
	return formGame, nil, http.StatusOK
}

func (r *AdPostgres) GetAllAd() ([]appl_row.AdFull, error, int) { // Получить все мероприятия за промежуток от сегодняшнего дня + 30 дней
	rowAds, err := r.db.Query("SELECT * FROM ad WHERE date_start > now() and date_start < now() + '30 days'::interval")
	if err != nil {
		return []appl_row.AdFull{}, fmt.Errorf("ошибка получения из базы данных, %s", err), http.StatusInternalServerError
	}
	defer rowAds.Close()
	var ad []appl_row.AdFull
	for rowAds.Next() {
		var p appl_row.AdFull
		if err := rowAds.Scan(
			&p.Id, &p.Title, &p.DateStart, &p.Created, &p.City,
			&p.Price, &p.Description, &p.EventType, &p.Participant, &p.SerialNumber,
			&p.PointOptions, &p.IsVisible, &p.IsFinished, &p.IsCancel, &p.LimitationTables,
		); err != nil {
			return []appl_row.AdFull{}, fmt.Errorf("ошибка преобразования полученных данных, %s", err), http.StatusInternalServerError
		}
		ad = append(ad, p)
	}

	return ad, nil, http.StatusOK
}

func (r *AdPostgres) ChangeLimitTable(adId int, newLimitationTables int) (error, int) { // Изменение количества допустимых столов с игроками в мероприятии
	if newLimitationTables < 1 {
		return fmt.Errorf("количество столов не может быть меньше 1"), http.StatusBadRequest
	} else {
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
				&p.PointOptions, &p.IsVisible, &p.IsFinished, &p.IsCancel, &p.LimitationTables,
			); err != nil {
				return fmt.Errorf("ошибка преобразования полученных данных, %s", err), http.StatusInternalServerError
			}
			ad = append(ad, p)
		}
		if len(ad) == 0 {
			return fmt.Errorf("данного мероприятия не существует"), http.StatusBadRequest
		} else {
			if ad[0].LimitationTables == newLimitationTables {
				return fmt.Errorf("значение не было изменено потому что старое и новое значения совпадают"), http.StatusBadRequest
			} else {
				_, err := r.db.Exec("UPDATE ad SET limitation_tables = $1 WHERE id = $2", newLimitationTables, adId)
				if err != nil {
					return fmt.Errorf("ошибка обновления из базы данных, %s", err), http.StatusInternalServerError
				}
			}
		}
	}
	return nil, http.StatusOK
}
