package service

import (
	httpapi "github.com/tsarkovmi/http_api"
	"github.com/tsarkovmi/http_api/pkg/repository"
)

//Позволяет генерировать код запустив инструмент который указа после команды
//используем установленную библиотеку мок, в которой указываем имя файла
//с интерфейсами и место куда сгенерировать моки
//go:generate mockgen -source=service.go -destination=mocks/mock.go

// Интерфейс для взаимодействия с основной логикой хэндлера
type CRUD interface {
	/*
		Метод, который принимает worker в качестве аргумента
		И возвращает id созданого в базе пользователя и ошибку
	*/
	CreateWorker(worker httpapi.Worker) (int, error)
	FindWorkerByID(workerId int) (httpapi.Worker, error)
	GetAllWorkers() ([]httpapi.Worker, error)
}

// Структура для интерфейса, вызывается из хэндлера
type Service struct {
	CRUD
}

// Сервис получает на вход Репозиторий, сам сервис вызывается из хэндлера
// ЭТО И ЕСТЬ ВНЕДРЕНИЕ ЗАВИСИМОСТЕЙ
func NewService(repos *repository.Repository) *Service {
	return &Service{
		CRUD: NewPostService(repos.CRUD),
	}
}
