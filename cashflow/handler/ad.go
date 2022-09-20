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
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err := h.services.AdChangeParams(body.City, body.Price)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Успешное изменение стоимости участия в городе %s", body.City),
	})
}

func (h *Handler) createAd(c *gin.Context) {
	type Body struct {
		Title         string `json:"title"`
		DateStart     string `json:"date_start"`
		City          string `json:"city"`
		Price         int    `json:"price"`
		Description   string `json:"description"`
		EventType     string `json:"event_type"`
		SerialNumber  int    `json:"serial_number"`
		PointsOptions int    `json:"points_options"`
		IsVisible     bool   `json:"is_visible"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err := h.services.CreateAd(appl_row.Ad(body))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Успешное создание мероприятия",
	})
}

func (h *Handler) activateAd(c *gin.Context) {
	type Body struct {
		Id int `json:"id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err := h.services.ActivateAd(body.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Успешная активация мероприятия",
	})
}

func (h *Handler) toCompleteAd(c *gin.Context) {
	type Body struct {
		Id int `json:"id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err := h.services.ToCompleteAd(body.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Успешное завершение мероприятия",
	})
}

func (h *Handler) cancelAd(c *gin.Context) {
	type Body struct {
		Id int `json:"id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err := h.services.CancelAd(body.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Успешная отмена мероприятия",
	})
}

func (h *Handler) entryToAd(c *gin.Context) {
	type Body struct {
		IdUser int `json:"id_user"`
		IdAd   int `json:"id_ad"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.EntryToAd(body.IdUser, body.IdAd)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  statusCode,
		"message": err,
	})
}

func (h *Handler) summarizingAd(c *gin.Context) {}
