package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	httpapi "github.com/tsarkovmi/http_api"
)

/*
type getAllListsResponse struct {
	Data []httpapi.Worker `json:"data"`
}
*/

// Создает JSON из фрагмента worker и записывает JSON в ответ
func (h *Handler) GetWorkers(c *gin.Context) {

	workers, err := h.services.CRUD.GetAllWorkers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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
func (h *Handler) PostWorkers(c *gin.Context) {
	var input httpapi.Worker

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
		//logrus.Errorf("error bindJSON: %s", err.Error())
	}

	/*
		ЗДЕСЬ ПОЛУЧИЛИ ДАННЫЕ И РАСПАРСИЛИ ИХ,
		ТЕПЕРЬ ДОЛЖНЫ ПЕРЕДАТЬ ДАННЫЕ НА СЛОЙ НИЖЕ
		В СЕРВИС
	*/

	id, err := h.services.CRUD.CreateWorker(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	/*
		//нужно добавлять в БД, а не в срез
		workers = append(workers, input) //в тупую добавляется в сущесвтующий срез
		c.IndentedJSON(http.StatusCreated, input)
	*/

}

/*
Поиск воркера по ID, цикл по срезу воркеров
При нахождении воркера - печатает его и статус
В ином случае воркер не найден
*/

func (h *Handler) GetWorkerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	worker, err := h.services.CRUD.FindWorkerByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, worker)
	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "workers not found"})

}
