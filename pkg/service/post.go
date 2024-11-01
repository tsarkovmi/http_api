package service

import (
	httpapi "github.com/tsarkovmi/http_api"
	"github.com/tsarkovmi/http_api/pkg/repository"
)

type PostService struct {
	repo repository.CRUD
}

// Конструктор для Service
func NewPostService(repo repository.CRUD) *PostService {
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
