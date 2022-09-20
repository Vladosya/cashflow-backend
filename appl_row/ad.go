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
