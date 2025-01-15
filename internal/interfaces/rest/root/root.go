package root

import (
	"meul/inventory/internal/infrastructures/drivers/postgres/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoot(r *gin.Engine, itemsDAO *models.ItemDAO) {
	r.GET("/", func(c *gin.Context) {
		// Retrieve all items from the database
		// items, err := itemsDAO.GetAllItems()
		// if err != nil {
		// 	// Handle error and return a 500 response
		// 	c.JSON(500, gin.H{
		// 		"error": "Failed to retrieve items",
		// 	})
		// 	return
		// }

		// Render the "index.html" template, passing the items as data by default
		// c.HTML(http.StatusOK, "index.html", gin.H{
		// 	"content": "views/items",
		// 	"items":   items,
		// })
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
