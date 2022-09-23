package appl_row

type SeatAtTable struct {
	Id     int `json:"id"`
	Name   any `json:"name"`
	IdUser any `json:"id_user"`
	Rank   any `json:"rank"`
}

type SeatAtTables struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	SeatAtTable []SeatAtTable `json:"seat_at_table"`
}

type Game struct {
	Id          int    `json:"id"`
	IdAd        int    `json:"id_ad"`
	SeatAtTable string `json:"seat_at_table"`
	IsFinished  bool   `json:"is_finished"`
}

type GameForm struct {
	Id          int            `json:"id"`
	IdAd        int            `json:"id_ad"`
	SeatAtTable []SeatAtTables `json:"seat_at_table"`
	IsFinished  bool           `json:"is_finished"`
}
