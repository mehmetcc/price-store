package db

import "time"

type PriceUpdate struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

func NewPriceUpdate(symbol string, price float64) *PriceUpdate {
	return &PriceUpdate{
		Symbol:    symbol,
		Price:     price,
		Timestamp: time.Now(),
	}
}
