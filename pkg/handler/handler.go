package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tsarkovmi/http_api/pkg/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/tsarkovmi/http_api/docs"
)

/*
Структура обработчика
*/
type Handler struct {
	services *service.Service
}

/*
Конструктор для структуры Handler. Инициализирует экземпляр структуры
Прослойка между хэндлером и сервисом
*/
func Newhandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

/*
Инициализирует и настраивает маршрутизатор gin
Устанавливаю gin в ReleaseMode, чтобы не логировалась отладочная информация
Далее устнанавливаются маршруты для эндпоинтов и для Swagger
Возвращает маршрутизатор с установленными маршрутами
*/
func (h *Handler) InitRourers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/workers", h.GetWorkers)
	router.GET("/workers/:id", h.GetWorkerByID)
	router.POST("/workers", h.PostWorkers)

	return router

}
