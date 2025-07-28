package main

import (
	"time"
	"warehouse-inventory/internal/db"
	"warehouse-inventory/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://warehouse-inventory-frontend.onrender.com", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.GET("/products", handler.GetProducts)
		api.POST("/products", handler.AddProduct)
		api.POST("/transaction", handler.AddTransaction)
		api.GET("/inventory", handler.GetInventory)
	}

	r.Run(":8080")
}
