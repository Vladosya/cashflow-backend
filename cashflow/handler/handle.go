package handler

import (
	"github.com/Vladosya/our_project/cashflow/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler { // создаём новый handler с полем services
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine { // обработчик роутов, Создание роутов
	router := gin.New() // инициализация роутов

	api := router.Group("/api-v1")
	{
		api.POST("/user/registration", h.registrationUser)
		api.POST("/ad/changeParams", h.adChangeParams)
	}

	return router
}
