package service

import "github.com/tsarkovmi/http_api/pkg/repository"

//Интерфейс для взаимодействия с основной логикой хэндлера
type CRUD interface {
}

//Структура для интерфейса, вызывается из хэндлера
type Service struct {
	CRUD
}

// Сервис получает на вход Репозиторий, сам сервис вызывается из хэндлера
// ЭТО И ЕСТЬ ВНЕДРЕНИЕ ЗАВИСИМОСТЕЙ
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
