package model

import "time"

type Product struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type StockMain struct {
	ID              uint      `gorm:"primaryKey"`
	TransactionType string    `json:"transactionType"`
	TransactionDate time.Time `json:"transactionDate"`
	Details         []StockDetail
}

type StockDetail struct {
	ID          uint `gorm:"primaryKey"`
	StockMainID uint
	ProductID   uint
	Product     Product `gorm:"foreignKey:ProductID"`
	Quantity    int
}
