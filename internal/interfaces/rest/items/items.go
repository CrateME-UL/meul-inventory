// Package items provides items rest api
package items

import (
	"meul/inventory/internal/infrastructures/drivers/postgres/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterItems(r *gin.Engine, itemsDAO *models.ItemDAO) {

	r.GET("/items", func(c *gin.Context) {
		items, err := itemsDAO.GetAllItems()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	})

	r.POST("/items", func(c *gin.Context) {
		var item models.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := itemsDAO.CreateItem(item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, item)
	})
}
