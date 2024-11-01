package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
Структура для респоса ошибки
*/
type errorResponse struct {
	Message string `json:"message"`
}

/*
type statusResponse struct {
	Status string `json:"status"`
}
*/

/*
Новый обработчик ошибок.
Логирует ошибки и отправляет клиенту сообщение с ошибкой в виде JSON
*/
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
