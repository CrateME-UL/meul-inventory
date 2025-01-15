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

		c.HTML(http.StatusOK, "items.html", gin.H{
			"Items": items,
		})
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
