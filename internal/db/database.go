package db

import (
	"log"
	"os"
	"warehouse-inventory/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	DB.AutoMigrate(&model.Product{}, &model.StockMain{}, &model.StockDetail{})
}
