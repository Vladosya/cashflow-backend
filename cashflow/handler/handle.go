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
		api.POST("/support/entryToAdSupport", h.entryToAdSupport) // Зарегистрировать пользователя на мероприятие бесплатно через службу поддержки
		api.POST("/user/registration", h.registrationUser)        // Регистрация пользователя
		api.POST("/user/entryToAd", h.entryToAd)                  // Вступление в мероприятие пользователем (пока без оплаты)
		api.POST("/user/refusalAd", h.refusalAd)                  // Отказ пользователя от мероприятия, в которое он уже вступил
		api.POST("/ad/changeParams", h.adChangeParams)            // Изменение стоимости участия мероприятия
		api.POST("/ad/createAd", h.createAd)                      // Создание мероприятия
		api.POST("/ad/activateAd", h.activateAd)                  // Активация мероприятия (сделать видимым мероприятие для всех пользователей)
		api.POST("/ad/toCompleteAd", h.toCompleteAd)              // Завершить мероприятия
		api.POST("/ad/cancelAd", h.cancelAd)                      // Отмена мероприятия (если мероприятие по каким-то причинам было отменено)
		api.POST("/ad/summarizingAd", h.summarizingAd)            // Распределение баллов по участникам за пройденное мероприятие
		api.GET("/ad/getAllAd", h.getAllAd)                       // Получить все мероприятия за промежуток от сегодняшнего дня + 30 дней
	}

	return router
}
