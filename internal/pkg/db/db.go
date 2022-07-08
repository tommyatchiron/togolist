package db

import (
	"fmt"

	"github.com/tommyatchiron/togolist/internal/pkg/config"
	"github.com/tommyatchiron/togolist/internal/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(config *config.Config, gormLogger *logger.GormLogger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Postgres.Dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("fail to initialize database")
	}
	return db, nil
}
