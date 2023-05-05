package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/wallet/app/usecase"
)

type DollarController struct{}

func NewDollarController() *DollarController {
	return &DollarController{}
}

func (d *DollarController) CurrentValue(c *gin.Context) {
	usecase := usecase.NewCurrentDollar()
	value, err := usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
	}

	c.JSON(http.StatusOK, map[string]any{"value": value})
}
