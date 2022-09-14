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
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err := h.services.RegistrationUser()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Успешное подтверждение профиля",
	})
}
