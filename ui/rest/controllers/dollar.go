package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DollarController struct{}

func NewDollarController() *DollarController {
	return &DollarController{}
}

func (d *DollarController) CurrentValue(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{"value": 1.23})
}
