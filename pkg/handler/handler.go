package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpapi "github.com/tsarkovmi/http_api"
)

type Handler struct {
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
}

// Создает JSON из фрагмента worker и записывает JSON в ответ
func GetWorkers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, workers)
}

/*
Тут должен быть response в котором будет
возвращаться новый уникальный ID
который здесь же и будет сгенерирован
*/

func PostWorkers(c *gin.Context) {
	var newWorker httpapi.Worker
	// вызов BindJSON чтобы привязать
	//полученный JSON к newWorker
	if err := c.BindJSON(&newWorker); err != nil {
		return
	}

	// добавление нового работника в срез
	workers = append(workers, newWorker)
	// возврат кода статуса
	c.IndentedJSON(http.StatusCreated, newWorker)

}

/*
Поиск воркера по ID, цикл по срезу воркеров
При нахождении воркера - печатает его и статус
В ином случае воркер не найден
*/
func GetWorkerByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range workers {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "workers not found"})

}
