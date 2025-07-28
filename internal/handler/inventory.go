package handler

import (
	"net/http"
	"warehouse-inventory/internal/db"

	"github.com/gin-gonic/gin"
)

type Inventory struct {
	ProductID uint   `json:"productId"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
}

func GetInventory(c *gin.Context) {
	var inventory []Inventory
	db.DB.Raw(`
		SELECT 
			p.id as product_id, p.name, 
			SUM(CASE sm.transaction_type WHEN 'IN' THEN sd.quantity ELSE 0 END) -
			SUM(CASE sm.transaction_type WHEN 'OUT' THEN sd.quantity ELSE 0 END) as quantity
		FROM stock_details sd
		JOIN stock_mains sm ON sm.id = sd.stock_main_id
		JOIN products p ON p.id = sd.product_id
		GROUP BY p.id, p.name
	`).Scan(&inventory)

	c.JSON(http.StatusOK, inventory)
}
