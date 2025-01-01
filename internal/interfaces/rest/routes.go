package rest

import (
	"meul/inventory/internal/interfaces/rest/health"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", health.PingHandler)
}
