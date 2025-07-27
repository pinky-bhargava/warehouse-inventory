package db

import (
	"log"
	"warehouse-inventory/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	DB.AutoMigrate(&model.Product{}, &model.StockMain{}, &model.StockDetail{})
}
