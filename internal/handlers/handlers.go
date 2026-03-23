package handlers

import (
	"fiberstarter/internal/repository"
	"fiberstarter/internal/service"
)

type Handlers struct {
	Service    *service.Service
	Repository *repository.Repository
}

func New(service *service.Service, repository *repository.Repository) *Handlers {
	return &Handlers{
		Service:    service,
		Repository: repository,
	}
}
