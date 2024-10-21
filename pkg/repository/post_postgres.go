package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	httpapi "github.com/tsarkovmi/http_api"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) CreateWorker(worker httpapi.Worker) (int, error) {
	var id int
	/*
		ВНИМАТЕЛЬНО, ЗДЕСЬ ЗАПИСЬ В БД И ЗАПРОС В БД
	*/
	query := fmt.Sprintf("INSERT INTO %s (name, age, salary, occupation) values ($1, $2, $3, $4) RETURNING id", workersTable)
	/*
		row хранит в себе информацию о возвращаемой строке из базы
		в данном случае - id
		Используя метод Scan запишем значение id в переменную,
		передав по ссылке
	*/
	row := r.db.QueryRow(query, worker.Name, worker.Age, worker.Salary, worker.Occupation)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
