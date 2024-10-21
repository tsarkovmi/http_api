package service

import (
	httpapi "github.com/tsarkovmi/http_api"
	"github.com/tsarkovmi/http_api/pkg/repository"
)

type PostService struct {
	repo repository.CRUD
}

func NewPostService(repo repository.CRUD) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreateWorker(worker httpapi.Worker) (int, error) {
	return s.repo.CreateWorker(worker)
}

func (s *PostService) FindWorkerByID(workerId int) (httpapi.Worker, error) {
	return s.repo.FindWorkerByID(workerId)
}
