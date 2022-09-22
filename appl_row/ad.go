package appl_row

import "time"

type Ad struct {
	Title            string `json:"title" binding:"required"`
	DateStart        string `json:"date_start" binding:"required"`
	City             string `json:"city" binding:"required"`
	Price            int    `json:"price" binding:"required"`
	Description      string `json:"description" binding:"required"`
	EventType        string `json:"event_type" binding:"required"`
	SerialNumber     int    `json:"serial_number" binding:"required"`
	PointsOptions    int    `json:"points_options" binding:"required"`
	IsVisible        bool   `json:"is_visible" binding:"required"`
	LimitationTables int    `json:"limitation_tables" binding:"required"`
}

type AdFull struct {
	Id               int       `json:"id"`
	Title            string    `json:"title"`
	DateStart        time.Time `json:"date_start"`
	Created          time.Time `json:"created"`
	City             string    `json:"city"`
	Price            int       `json:"price"`
	Description      string    `json:"description"`
	EventType        string    `json:"ок"`
	Participant      []uint8   `json:"participant"`
	SerialNumber     int       `json:"serial_number"`
	PointOptions     int       `json:"point_options"`
	IsVisible        bool      `json:"is_visible"`
	IsFinished       bool      `json:"is_finished"`
	IsCancel         bool      `json:"is_cancel"`
	LimitationTables int       `json:"limitation_tables"`
}

type WinUser struct {
	Id       int `json:"id"`
	Place    int `json:"place"`
	Assigned int `json:"assigned"`
}

type WinnersPart struct {
	Name    string    `json:"name"`
	WinUser []WinUser `json:"winUser"`
}

type WinRes struct {
	Id           int `json:"id"`
	Place        int `json:"place"`
	NumberPoints int `json:"numberPoints"`
}

type Scoring struct {
	WinRes []WinRes `json:"winRes"`
}

type PointsGame struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	City    string  `json:"city"`
	Version int     `json:"version"`
	Scoring []uint8 `json:"scoring"`
}
