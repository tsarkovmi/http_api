package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tsarkovmi/http_api/pkg/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/tsarkovmi/http_api/docs"
)

type Handler struct {
	services *service.Service
}

func Newhandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRourers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/workers", h.GetWorkers)
	router.GET("/workers/:id", h.GetWorkerByID)
	router.POST("/workers", h.PostWorkers)

	return router

}
