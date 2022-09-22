package handler

import (
	"fmt"
	"github.com/Vladosya/our_project/appl_row"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) adChangeParams(c *gin.Context) {
	type Body struct {
		City  string `json:"city"`
		Price int    `json:"price"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.AdChangeParams(body.City, body.Price)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": fmt.Sprintf("успешное изменение стоимости участия в городе %s", body.City),
	})
}

func (h *Handler) createAd(c *gin.Context) {
	type Body struct {
		Title            string `json:"title"`
		DateStart        string `json:"date_start"`
		City             string `json:"city"`
		Price            int    `json:"price"`
		Description      string `json:"description"`
		EventType        string `json:"event_type"`
		SerialNumber     int    `json:"serial_number"`
		PointsOptions    int    `json:"points_options"`
		IsVisible        bool   `json:"is_visible"`
		LimitationTables int    `json:"limitation_tables"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.CreateAd(appl_row.Ad(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное создание мероприятия",
	})
}

func (h *Handler) activateAd(c *gin.Context) {
	type Body struct {
		Id int `json:"id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.ActivateAd(body.Id)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная активация мероприятия",
	})
}

func (h *Handler) toCompleteAd(c *gin.Context) {
	type Body struct {
		Id int `json:"id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.ToCompleteAd(body.Id)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное завершение мероприятия",
	})
}

func (h *Handler) cancelAd(c *gin.Context) {
	type Body struct {
		Id int `json:"id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.CancelAd(body.Id)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная отмена мероприятия",
	})
}

func (h *Handler) summarizingAd(c *gin.Context) {
	type Body struct {
		IdAd        int                    `json:"id_ad"`
		WinnersPart []appl_row.WinnersPart `json:"winners_part"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.SummarizingAd(body.IdAd, body.WinnersPart)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное распределение баллов",
	})
}

func (h *Handler) replantAd(c *gin.Context) {
	type Body struct {
		IdAd         int                     `json:"id_ad"`
		SeatAtTables []appl_row.SeatAtTables `json:"seat_at_tables"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.ReplantAd(body.IdAd, body.SeatAtTables)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное распределение пользователей по столам",
	})
}

func (h *Handler) getAllAd(c *gin.Context) {
	res, err, statusCode := h.services.GetAllAd()
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  res,
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное получение всех мероприятий",
		"result":  res,
	})
}
