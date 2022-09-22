package repository

import (
	"encoding/json"
	"fmt"
	"github.com/Vladosya/our_project/helpers"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type SupportPostgres struct {
	db *sqlx.DB
}

func NewSupportPostgres(db *sqlx.DB) *SupportPostgres {
	return &SupportPostgres{
		db: db,
	}
}

func (r *SupportPostgres) EntryToAdSupport(userId int, adId int) (error, int) {
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
		var arr []int
		err := json.Unmarshal(ad[0].Participant, &arr)
		if err != nil {
			return fmt.Errorf("ошибка работы с json, %s", err), http.StatusInternalServerError
		}
		if helpers.ContainsInt(arr, userId) == true {
			return fmt.Errorf("данный пользователь уже зарегистрирован на данное мероприятие"), http.StatusBadRequest
		} else {
			var resCalculateByTableLimit = helpers.CalculateByTableLimit(ad[0].LimitationTables, len(arr))
			if resCalculateByTableLimit == false {
				return fmt.Errorf("данный пользователь не можете вступить в данное мероприятие, т.к оно полностью заполнено"), http.StatusBadRequest
			} else {
				_, err = r.db.Exec("UPDATE ad SET participant = participant || $1 WHERE id = $2", userId, adId)
				if err != nil {
					return fmt.Errorf("ошибка обновления из базы данных, %s", err), http.StatusInternalServerError
				}
			}
		}
	}
	return nil, http.StatusOK
}
