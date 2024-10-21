package repository

import (
	"github.com/jmoiron/sqlx"
	httpapi "github.com/tsarkovmi/http_api"
)

// Интерфейс для взаимодействия с основной логикой хэндлеря
type CRUD interface {
	CreateWorker(worker httpapi.Worker) (int, error)
}

// Структура для интерфейса
type Repository struct {
	CRUD
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CRUD: NewPostPostgres(db),
	}
}
