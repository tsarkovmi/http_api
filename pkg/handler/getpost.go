package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	httpapi "github.com/tsarkovmi/http_api"
)

//	@Summary		GetWorkers
//	@Tags			POST
//	@Description	Post Worker to DB
//	@ID				get-workers
//	@Accept			json
//	@Produce		json
//  @Param 			input json httpapi.Worker
//	@Success		200		{integer}	integer	1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/workers/ [post]

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

//	@Summary		GetWorkers
//	@Tags			GET
//	@Description	Get All workers from DB
//	@ID				get-workers
//	@Accept			json
//	@Produce		json
//	@Success		200		{integer}	integer	1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/workers/ [get]

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

//	@Summary		GetWorkerByID
//	@Tags			GET
//	@Description	Get One workerr from DB
//	@ID				get-worker
//	@Accept			json
//	@Produce		json
//  @Param 			id json int
//	@Success		200		{integer}	integer	1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
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
