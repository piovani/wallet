package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
