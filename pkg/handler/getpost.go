package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	httpapi "github.com/tsarkovmi/http_api"
)

// Создает JSON из фрагмента worker и записывает JSON в ответ
func GetWorkers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, workers)
}

/*
Тут должен быть response в котором будет
возвращаться новый уникальный ID
который здесь же и будет сгенерирован

вызов BindJSON чтобы привязать
полученный JSON к newWorker
добавление нового работника в срез
возврат кода статуса
*/
func PostWorkers(c *gin.Context) {
	var newWorker httpapi.Worker

	if err := c.BindJSON(&newWorker); err != nil {
		logrus.Errorf("error bindJSON: %s", err.Error())
	}
	workers = append(workers, newWorker)
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
