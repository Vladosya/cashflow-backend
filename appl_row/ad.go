package appl_row

type Ad struct {
	Title         string `json:"title" binding:"required"`
	DateStart     string `json:"date_start" binding:"required"`
	City          string `json:"city" binding:"required"`
	Price         int    `json:"price" binding:"required"`
	Description   string `json:"description" binding:"required"`
	EventType     string `json:"event_type" binding:"required"`
	SerialNumber  int    `json:"serial_number" binding:"required"`
	PointsOptions int    `json:"points_options" binding:"required"`
	IsVisible     bool   `json:"is_visible" binding:"required"`
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
