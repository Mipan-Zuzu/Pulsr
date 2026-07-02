package database

import (
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnection () (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASAE_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}