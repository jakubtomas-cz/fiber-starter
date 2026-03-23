package service

import (
	"fiberstarter/internal/repository"
	"fiberstarter/internal/utils"
)

type Service struct {
	Repository *repository.Repository
}

func New(repository *repository.Repository) (*Service, error) {
	if repository == nil {
		return nil, utils.ErrNoRepository
	}

	service := &Service{
		Repository: repository,
	}

	return service, nil
}
