package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	httpapi "github.com/tsarkovmi/http_api"
)

// PostWorkers godoc
//	@Summary		Add a new worker
//	@Description	Создание нового работника в базе данных
//	@Tags			workers
//	@Accept			json
//	@Produce		json
//	@Param			input	body		httpapi.Worker			true	"Worker info"
//	@Success		200		{object}	map[string]interface{}	"ID созданного работника"
//	@Failure		400		{object}	ErrorResponse			"Неверный формат данных"
//	@Failure		500		{object}	ErrorResponse			"Ошибка сервера"
//	@Router			/workers [post]

func (h *Handler) PostWorkers(c *gin.Context) {
	var input httpapi.Worker

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
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
}

// GetWorkers godoc
//	@Summary		Get all workers
//	@Description	Получение списка всех работников из базы данных
//	@Tags			workers
//	@Produce		json
//	@Success		200	{object}	serResponse		"Список всех работников"
//	@Failure		500	{object}	ErrorResponse	"Ошибка сервера"
//	@Router			/workers [get]

func (h *Handler) GetWorkers(c *gin.Context) {

	workers, err := h.services.CRUD.GetAllWorkers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := initSerResponse()

	err = resp.serializeWorker(workers)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, resp)

}

// GetWorkerByID godoc
//	@Summary		Get a worker by ID
//	@Description	Получение информации о работнике по его ID
//	@Tags			workers
//	@Produce		json
//	@Param			id	path		int				true	"Worker ID"
//	@Success		200	{object}	serResponse		"Данные о работнике"
//	@Failure		500	{object}	ErrorResponse	"Ошибка сервера"
//	@Failure		400	{object}	ErrorResponse	"Неверный ID"
//	@Router			/workers/{id} [get]

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

	resp := initSerResponse()
	err = resp.serializeWorker(worker)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, resp)

}
