package pages

import (
	"fiberstarter/internal/repository"
	"fiberstarter/internal/service"
)

type Pages struct {
	Service    *service.Service
	Repository *repository.Repository
}

func New(service *service.Service, repository *repository.Repository) *Pages {
	return &Pages{
		Service:    service,
		Repository: repository,
	}
}
