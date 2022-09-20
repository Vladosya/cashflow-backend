package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) registrationUser(c *gin.Context) {
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
	err, statusCode := h.services.RegistrationUser()
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": err.Error(),
	})
}

func (h *Handler) entryToAd(c *gin.Context) {
	type Body struct {
		IdUser int `json:"id_user"`
		IdAd   int `json:"id_ad"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.EntryToAd(body.IdUser, body.IdAd)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": err.Error(),
	})
}
