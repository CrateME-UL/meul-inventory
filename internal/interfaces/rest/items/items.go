// Package items provides items rest api
package items

import (
	"meul/inventory/internal/infrastructures/drivers/postgres/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterItems(r *gin.Engine, itemsDAO *models.ItemDAO) {

	r.GET("/items", func(c *gin.Context) {
		items, err := itemsDAO.GetAllItems()
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": err.Error(),
			})
			return
		}

		// Check if the request is an HTMX request
		if c.GetHeader("HX-Request") != "" {
			c.HTML(http.StatusOK, "views/items", items)
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"content": "views/items",
				"items":   items,
			})
		}
	})

	r.POST("/items", func(c *gin.Context) {
		name := c.PostForm("name")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Item name is required"})
			return
		}

		item := models.Item{ItemNumber: uuid.New(), Name: name}

		if err := itemsDAO.CreateItem(item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusCreated, "components/item", item)
	})

}
