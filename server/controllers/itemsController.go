package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dasdachs/inventory/models"
)

func GetAllItems(c *gin.Context) {
	items, err := models.GetAllItems()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

func CreateItem(c *gin.Context) {
	item := new(models.Item)

	if err := c.ShouldBindJSON(item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := item.Save(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusCreated, item)
}

func UpdateItem(c *gin.Context) {
	id := c.Params.ByName("id")

	item, err := models.GetItemById(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var updatedItem models.Item

	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(updatedItem.ItemName) != 0 {
		(*item).ItemName = updatedItem.ItemName
	}

	(*item).Quantity = updatedItem.Quantity

	if err := item.Update(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusCreated, item)
	}
}

func GetItemById(c *gin.Context) {
	id := c.Params.ByName("id")

	if item, err := models.GetItemById(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, item)
	}
}

func DeleteItem(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := models.DeleteItem(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}
