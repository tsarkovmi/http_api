package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	workersTable = "workers"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

/*
Создаем указатель на базу данных, открывая её с помощью полей Config
Ниже проверяем есть ли соединение с БД и возвращаем указатель на БД
*/
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
