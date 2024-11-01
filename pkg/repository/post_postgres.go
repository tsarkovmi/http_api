package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	httpapi "github.com/tsarkovmi/http_api"
)

type PostPostgres struct {
	db *sqlx.DB
}

/*
Конструктор для создания указателя на оболочку БД
*/
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

func (r *PostPostgres) FindWorkerByID(workerId int) (httpapi.Worker, error) {
	var worker httpapi.Worker
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", workersTable)
	err := r.db.Get(&worker, query, workerId)

	return worker, err
}

func (r *PostPostgres) GetAllWorkers() ([]httpapi.Worker, error) {
	var workers []httpapi.Worker
	query := fmt.Sprintf("SELECT * FROM %s", workersTable)
	err := r.db.Select(&workers, query)
	return workers, err
}

//func (r *PostPostgres) AllWorkers() (httpapi.Worker)
