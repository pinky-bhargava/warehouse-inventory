package main

import (
	"warehouse-inventory/internal/db"
	"warehouse-inventory/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/products", handler.GetProducts)
		api.POST("/products", handler.AddProduct)
		api.POST("/transaction", handler.AddTransaction)
		api.GET("/inventory", handler.GetInventory)
	}

	r.Run(":8080")
}
