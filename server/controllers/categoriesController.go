package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dasdachs/inventory/models"
)

func GetAllCategories(c *gin.Context) {
	categories, err := models.GetAllCategories()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, categories)
	}
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := category.SaveCategory(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusCreated, &category)
}

func UpdateCategory(c *gin.Context) {
	id := c.Params.ByName("id")

	category, err := models.GetCategoryById(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var updatedCategory models.Category

	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(updatedCategory.CategoryName) != 0 {
		(*category).CategoryName = updatedCategory.CategoryName
	}

	if err := category.UpdateCategory(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusCreated, category)
	}
}

func AppendItemToCategory(c *gin.Context) {
	id := c.Params.ByName("id")

	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := item.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if appendErr := models.AppendItem(id, &item); appendErr != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, &item)
	}
}

func GetCategoryById(c *gin.Context) {
	id := c.Params.ByName("id")

	category, err := models.GetCategoryById(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		if category.ID == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, category)
		}
	}
}

func DeleteCategory(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := models.DeleteCategory(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}
