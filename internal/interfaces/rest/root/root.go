package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoot(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
