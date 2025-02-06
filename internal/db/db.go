package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(dsn string) {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	if err := db.AutoMigrate(&PriceUpdate{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}
}

func Create(pu *PriceUpdate) error {
	result := db.Create(pu)
	return result.Error
}
