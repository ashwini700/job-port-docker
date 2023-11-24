package database

import (
	"job-port-api/config"
	"job-port-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(cfg config.Config) (*gorm.DB, error) {
	// dataSources := "host=postgres user=postgres password=Ashwini dbname=postgres port=5432 "
	dsn := cfg.DatabaseCOnfig.DB_DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Job{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
