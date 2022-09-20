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
		api.POST("/user/registration", h.registrationUser) // Регистрация пользователя
		api.POST("/ad/changeParams", h.adChangeParams)     // Изменение стоимости участия мероприятия
		api.POST("/ad/createAd", h.createAd)               // Создание мероприятия
		api.POST("/ad/activateAd", h.activateAd)           // Активация мероприятия (сделать видимым мероприятие для всех пользователей)
		api.POST("/ad/toCompleteAd", h.ToCompleteAd)       // Завершить мероприятия
		api.POST("/ad/summarizingAd", h.summarizingAd)     // распределение баллов по участникам за пройденное мероприятие
	}

	return router
}
