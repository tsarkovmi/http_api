package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type worker struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Age        int16   `json:"age"`
	Salary     float32 `json:"salary"`
	Occupation string  `json:"occupation"`
}

// Создает JSON из фрагмента worker и записывает JSON в ответ
func getWorkers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, workers)
}

/*
Тут должен быть response в котором будет
возвращаться новый уникальный ID
который здесь же и будет сгенерирован
*/

func postWorkers(c *gin.Context) {
	var newWorker worker

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

var workers = []worker{
	{ID: "1", Name: "Mike Vazov", Age: 46, Salary: 450123.23, Occupation: "Переводчик"},
	{ID: "2", Name: "Nikolay Mazurin", Age: 24, Salary: 320123.23, Occupation: "Сварщик"},
	{ID: "3", Name: "Alexey Popov", Age: 64, Salary: 120123.56, Occupation: "Руководитель группы"},
}

/*
Поиск воркера по ID, цикл по срезу воркеров
При нахождении воркера - печатает его и статус
В ином случае воркер не найден
*/
func getWorkerByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range workers {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "workers not found"})

}

func main() {
	/*
		Инициализировали router

	*/
	router := gin.Default()
	router.GET("/workers", getWorkers)
	router.GET("/workers/:id", getWorkerByID)
	router.POST("/workers", postWorkers)

	router.Run("localhost:8080")

}
