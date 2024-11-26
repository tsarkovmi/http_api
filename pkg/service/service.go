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

/*НАШЁЛ*/
type PostService struct {
	repo repository.CRUD
}

// Конструктор для Service
func NewService(repo repository.CRUD) *PostService {
	return &PostService{repo: repo}
}

// Реализация метода CreateWorker для интерфейса на слое Service
func (s *PostService) CreateWorker(worker httpapi.Worker) (int, error) {
	return s.repo.CreateWorker(worker)
}

// Реализация метода FindWorkerByID для интерфейса на слое Service
func (s *PostService) FindWorkerByID(workerId int) (httpapi.Worker, error) {
	return s.repo.FindWorkerByID(workerId)
}

// Реализация метода GetAllWorkers для интерфейса на слое Service
func (s *PostService) GetAllWorkers() ([]httpapi.Worker, error) {
	return s.repo.GetAllWorkers()
}
