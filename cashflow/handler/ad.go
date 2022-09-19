package handler

import (
	"fmt"
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
