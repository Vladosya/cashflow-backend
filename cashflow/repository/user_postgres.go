package repository

import (
	"fmt"
	"github.com/Vladosya/our_project/helpers"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

type Ad struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	DateStart    time.Time `json:"date_start"`
	Created      time.Time `json:"created"`
	City         string    `json:"city"`
	Price        int       `json:"price"`
	Description  string    `json:"description"`
	EventType    string    `json:"ок"`
	Participant  []uint8   `json:"participant"`
	SerialNumber int       `json:"serial_number"`
	PointOptions int       `json:"point_options"`
	NumOfTables  int       `json:"num_of_tables"`
	IsVisible    bool      `json:"is_visible"`
	IsFinished   bool      `json:"is_finished"`
	IsCancel     bool      `json:"is_cancel"`
}

func (r *UserPostgres) RegistrationUser() (error, int) {
	fmt.Println("RegistrationUser")
	fmt.Println("helpers.RandomStrGeneration(8)", helpers.RandomStrGeneration(8))
	return fmt.Errorf("успешная регистрация пользователя"), http.StatusOK
}

func (r *UserPostgres) EntryToAd(userId int, adId int) (error, int) {
	rowAd, err := r.db.Query("SELECT * FROM ad WHERE id = $1 AND $2 = any(participant)", adId, userId)
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
			&p.PointOptions, &p.NumOfTables, &p.IsVisible, &p.IsFinished, &p.IsCancel,
		); err != nil {
			return fmt.Errorf("ошибка преобразования полученных данных, %s", err), http.StatusInternalServerError
		}
		ad = append(ad, p)
	}
	if len(ad) == 1 {
		return fmt.Errorf("вы уже зарегистрированы на данное мероприятие"), http.StatusBadRequest
	} else {
		_, err := r.db.Exec("UPDATE ad SET participant = participant || $1::INTEGER WHERE id = $2", userId, adId)
		if err != nil {
			return fmt.Errorf("ошибка изменения из базы данных, %s", err), http.StatusInternalServerError
		}
	}
	return fmt.Errorf("успешное вступление в мероприятие"), http.StatusOK
}
