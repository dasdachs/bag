package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dasdachs/inventory/models"
)

func GetAllItems(c *gin.Context) {
	items, err := models.GetAllItems()

	if err != nil {
		c.AbortWithStatus(500)
	} else {
		c.JSON(200, items)
	}
}

func CreateOrUpdateItem(c *gin.Context) {
	item := new(models.Item)

	if err := c.ShouldBindJSON(item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := item.Save(); err != nil {
		c.AbortWithStatus(500)
	}

	c.JSON(201, item)
}

func GetItemById(c *gin.Context) {
	id := c.Params.ByName("id")

	if item, err := models.GetItemById(id); err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, item)
	}
}

func DeleteItem(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := models.DeleteItem(id); err != nil {
		c.AbortWithStatus(404)
	} else {
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}
