package main

import (
	"fiberstarter/internal/repository"
	"fiberstarter/internal/service"
	"fiberstarter/internal/transport"
	"fiberstarter/internal/utils"
	"log"
)

func init() {
	if err := utils.Init(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	repository, err := repository.New()
	if err != nil {
		log.Fatal(err)
	}

	service, err := service.New(repository)
	if err != nil {
		log.Fatal(err)
	}

	if err := transport.New(service, repository); err != nil {
		log.Fatal(err)
	}
}
