package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/piovani/wallet/infra/config"
)

type Rest struct {
	http *gin.Engine
}

func NewRest() *Rest {
	rest := &Rest{
		http: gin.Default(),
	}
	GetRoute(rest.http)
	return rest
}

func (r *Rest) Start() error {
	return r.http.Run(fmt.Sprintf(":%d", config.Env.ApiRestPort))
}
