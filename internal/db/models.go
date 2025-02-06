package db

import "time"

type PriceUpdate struct {
	ID        uint `gorm:"primaryKey"`
	Symbol    string
	Price     float64
	Timestamp time.Time
}
