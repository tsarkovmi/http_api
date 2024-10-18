package handler

import (
	"github.com/gin-gonic/gin"
	httpapi "github.com/tsarkovmi/http_api"
	"github.com/tsarkovmi/http_api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func Newhandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRourers() *gin.Engine {
	router := gin.New()

	router.GET("/workers", GetWorkers)
	router.GET("/workers/:id", GetWorkerByID)
	router.POST("/workers", PostWorkers)

	return router

}

var workers = []httpapi.Worker{
	{ID: "1", Name: "Mike Vazov", Age: 46, Salary: 450123.23, Occupation: "Переводчик"},
	{ID: "2", Name: "Nikolay Mazurin", Age: 24, Salary: 320123.23, Occupation: "Сварщик"},
	{ID: "3", Name: "Alexey Popov", Age: 64, Salary: 120123.56, Occupation: "Руководитель группы"},
	{ID: "4", Name: "Anna Leonova", Age: 18, Salary: 170673.56, Occupation: "MANAGER"},
}
