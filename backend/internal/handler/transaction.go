package handler

import (
	"net/http"
	"time"
	"warehouse-inventory/internal/db"
	"warehouse-inventory/internal/model"

	"github.com/gin-gonic/gin"
)

type TransactionItem struct {
	ProductID uint `json:"productId"`
	Quantity  int  `json:"quantity"`
}

type TransactionRequest struct {
	TransactionType string            `json:"transactionType"`
	TransactionDate time.Time         `json:"transactionDate"`
	Items           []TransactionItem `json:"items"`
}

func AddTransaction(c *gin.Context) {
	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	main := model.StockMain{
		TransactionType: req.TransactionType,
		TransactionDate: req.TransactionDate,
	}
	db.DB.Create(&main)

	for _, item := range req.Items {
		db.DB.Create(&model.StockDetail{
			StockMainID: main.ID,
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
		})
	}

	c.Status(http.StatusOK)
}
