package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	httpapi "github.com/tsarkovmi/http_api"
)

func (h *Handler) GetWorkers(c *gin.Context) {

	workers, err := h.services.CRUD.GetAllWorkers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := initSerResponse()
	for _, a := range workers {
		err := resp.serializeWorker(a)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
	}
	c.IndentedJSON(http.StatusOK, resp)

}

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
	}
	c.IndentedJSON(http.StatusOK, resp)

}
