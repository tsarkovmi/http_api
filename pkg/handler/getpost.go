package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	httpapi "github.com/tsarkovmi/http_api"
)

/*
	Обработка POST запроса для создания нового worker в БД
	Извлекает и валидирует JSON из тела запроса
	Если входные данные некорректны метод возвращает HTTP ответ с кодом 400
	и сообщением об ошибке
	Далее данные передаеются на слой сервиса через метод CreateWorker
	В случае ошибки на сервисе - 500 и сообщение об ошибке
	Если всё ок, то возвращает ID созданной записи и 200
*/
// PostWorkers godoc
//
//	@Summary		Add a new worker
//	@Description	Этот метод обрабатывает POST-запрос для создания нового воркера
//	@Tags			workers
//	@Accept			json
//	@Produce		json
//	@Param			input	body		httpapi.Worker			true	"Worker info"
//	@Success		200		{object}	map[string]interface{}	"ID созданного работника"
//	@Failure		400		{object}	errorResponse			"Неверный формат данных"
//	@Failure		500		{object}	errorResponse			"Ошибка сервера"
//	@Router			/workers [post]
func (h *Handler) PostWorkers(c *gin.Context) {
	var input httpapi.Worker

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
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

/*
	GET запрос для получения списка всех worker
	Вызывает метод GetAllWokers из слоя сервиса, чтобы получить все записи
	В случае ошибки - 500 и сообщение об ошибки
	Далее данные передаются в метод serializeWorker для форматирования
	Возвращает JSON со всеми worker в 3 форматах, и кодом 200
*/
// GetWorkers godoc
//
//	@Summary		Get all workers
//	@Description	Получение списка всех worker из базы данных
//	@Tags			workers
//	@Produce		json
//	@Success		200	{object}	serResponse		"Список всех работников"
//	@Failure		500	{object}	errorResponse	"Ошибка сервера"
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

/*
	GET запрос для получения конкретного worker по ID
	Извлечение ID с помощью Atoi и проверка валидации
	Далее вызывается метод из слоя сервиса, для получения данных
	Если рабочий не найден 404 и сообщение об ошибке
	В случае успеха данные передаются для форматирования в serializeWorker
	Возвращает JSON с тремя полями в заданных форматах
*/
// GetWorkerByID godoc
//
//	@Summary		Get a worker by ID
//	@Description	Получение информации о работнике по его ID
//	@Tags			workers
//	@Produce		json
//	@Param			id	path		int				true	"Worker ID"
//	@Success		200	{object}	serResponse		"Данные о работнике"
//	@Failure		500	{object}	errorResponse	"Ошибка сервера"
//	@Failure		400	{object}	errorResponse	"Неверный ID"
//	@Failure		404	{object}	errorResponse	"Работник не найден"
//	@Router			/workers/{id} [get]
func (h *Handler) GetWorkerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid worker ID")
		return
	}

	worker, err := h.services.CRUD.FindWorkerByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "worker not found")
		return
	}

	resp := initSerResponse()
	err = resp.serializeWorker(worker)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, resp)

}
