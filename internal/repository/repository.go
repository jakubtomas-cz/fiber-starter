package repository

import (
	"context"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	DB  *gorm.DB
	ctx context.Context
	mu  sync.Mutex
}

func New() (*Repository, error) {
	repository := &Repository{
		ctx: context.Background(),
	}

	db, err := gorm.Open(repository.GetDialect(), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return repository, err
	}

	repository.DB = db

	// TODO: Uncomment to enable auto migrations from models
	// err = db.AutoMigrate()
	// if err != nil {
	// 	return repository, err
	// }

	return repository, nil
}
