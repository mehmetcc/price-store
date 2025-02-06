package db

import "time"

type PriceUpdate struct {
	ID        uint `gorm:"primaryKey"`
	Symbol    string
	Price     float64
	Timestamp time.Time
}

func NewPriceUpdate(symbol string, price float64) *PriceUpdate {
	return &PriceUpdate{
		Symbol:    symbol,
		Price:     price,
		Timestamp: time.Now(),
	}
}
