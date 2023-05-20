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

func (d *DollarController) PurchaseValues(c *gin.Context) {
	values, err := dollar.NewPurchaseValues().Execute()
	PurchaseValuesOut := NewPurchaseValuesOut(values)
	if len(err) > 0 {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}
	c.JSON(http.StatusOK, PurchaseValuesOut)
}
