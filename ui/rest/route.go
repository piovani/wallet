package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/piovani/wallet/ui/rest/controllers"
)

func GetRoute(r *gin.Engine) {
	// Controllers
	healthController := controllers.NewHealthController()
	dollarController := controllers.NewDollarController()

	// HEALTH
	r.GET("/health", healthController.Health)

	// DOLLAR
	r.GET("dollar/current-value", dollarController.CurrentValue)
}
