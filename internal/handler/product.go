package handler

import (
	"net/http"
	"warehouse-inventory/internal/db"
	"warehouse-inventory/internal/model"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []model.Product
	db.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context) {
	var p model.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&p)
	c.JSON(http.StatusCreated, p)
}
