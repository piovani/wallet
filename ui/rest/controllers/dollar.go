package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/wallet/app/usecase/dollar"
)

type DollarController struct{}

func NewDollarController() *DollarController {
	return &DollarController{}
}

// func (d *DollarController) CurrentValue(c *gin.Context) {
// 	// usecase := dollar.NewCurrentDollar()
// 	// value, err := usecase.Execute()
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
// 	// }

// 	c.JSON(http.StatusOK, map[string]any{"value": value})
// }

func (d *DollarController) PurchaseValues(c *gin.Context) {
	values, err := dollar.NewPurchaseValues().Execute()
	if len(err) > 0 {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, values)
}
