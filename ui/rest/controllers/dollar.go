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
	if len(err) > 0 {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, values)
}
