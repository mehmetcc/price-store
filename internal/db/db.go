package db

import (
	"errors"
	"log"

	"github.com/mehmetcc/symbol-store/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(cfg *config.Config) {
	var err error
	db, err = gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	if err := db.AutoMigrate(&PriceUpdate{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}
}

func Create(pu *PriceUpdate) error {
	if pu.Price == 0 {
		return errors.New("price cannot be 0")
	}
	result := db.Create(pu)
	return result.Error
}

func GetPriceUpdates(page, pageSize int) ([]PriceUpdate, error) {
	var priceUpdates []PriceUpdate
	offset := (page - 1) * pageSize
	result := db.Offset(offset).Limit(pageSize).Find(&priceUpdates)
	if result.Error != nil {
		return nil, result.Error
	}

	return priceUpdates, nil
}

func GetTotalPriceUpdatesCount() (int64, error) {
	var count int64
	result := db.Model(&PriceUpdate{}).Count(&count)
	return count, result.Error
}

func SearchPriceUpdatesBySymbol(symbol string, page, pageSize int) ([]PriceUpdate, error) {
	var priceUpdates []PriceUpdate
	offset := (page - 1) * pageSize
	result := db.Where("symbol = ?", symbol).
		Offset(offset).
		Limit(pageSize).
		Find(&priceUpdates)
	if result.Error != nil {
		return nil, result.Error
	}
	return priceUpdates, nil
}

func GetFilteredPriceUpdatesCount(symbol string) (int64, error) {
	var count int64
	result := db.Model(&PriceUpdate{}).Where("symbol = ?", symbol).Count(&count)
	return count, result.Error
}
