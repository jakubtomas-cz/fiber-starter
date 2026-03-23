package repository

import (
	"fiberstarter/internal/utils"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (r *Repository) GetDialect() gorm.Dialector {
	databaseURL := utils.ENV.DatabaseURL

	if strings.HasPrefix(databaseURL, "postgres") {
		return postgres.Open(databaseURL)
	}

	return sqlite.Open(databaseURL)
}
