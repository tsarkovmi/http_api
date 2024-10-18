package repository

import "github.com/jmoiron/sqlx"

//Интерфейс для взаимодействия с основной логикой хэндлеря
type CRUD interface {
}

//Структура для интерфейса
type Repository struct {
	CRUD
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
