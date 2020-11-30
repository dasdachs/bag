package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm/clause"
)

type Category struct {
	Id    uint   `gorm:"primaryKey autoIncrement" form:"id" json:"id"`
	Name  string `gorm:"not null unique" form:"name" json:"name"`
	Items []Item `form:"category" json:"items"`
}

type Item struct {
	Id         uint     `gorm:"primaryKey autoIncrement" form:"id" json:"id"`
	Name       string   `gorm:"not null" form:"name" json:"name"`
	Quantity   uint8    `gorm:"not null" form:"quantity" json:"quantity"`
	CategoryID uint     `form:"categoryId" json:"categoryId"`
	Category   Category `form:"category" json:"category"`
}

func main() {
	db, err := gorm.Open("sqlite3", "./inventory.db")
	if err != nil {
		fmt.Println(err)
	}
	db.LogMode(true)

	defer db.Close()

	db.AutoMigrate(&Category{}, &Item{})

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	router.GET("/api/v1/items", func(c *gin.Context) {
		var items []Item

		if err := db.Preload("Category").Find(&items).Error; err != nil {
			c.AbortWithStatus(500)
		} else {
			c.JSON(200, items)
		}
	})

	router.POST("/api/v1/items", func(c *gin.Context) {
		var item Item

		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err := db.Create(&item).Error; err != nil {
			c.AbortWithStatus(500)
		}

		c.JSON(201, item)
	})

	router.GET("/api/v1/items/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var item Item

		if err := db.Preload("Category").Where("id = ?", id).First(&item).Error; err != nil {
			c.AbortWithStatus(404)
		} else {

			c.JSON(200, item)
		}
	})

	router.PUT("/api/v1/items/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var item Item

		if err := db.Where("id = ?", id).First(&item).Error; err != nil {
			c.AbortWithStatus(404)
		} else {
			db.Save(&item)
			c.JSON(200, item)
		}
	})

	router.DELETE("/api/v1/items/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var item Item

		if err := db.Where("id = ?", id).First(&item).Error; err != nil {
			c.AbortWithStatus(404)
		} else {
			db.Delete(&item)
			c.Writer.WriteHeader(http.StatusNoContent)
		}
	})

	// Categories

	router.GET("/api/v1/categories", func(c *gin.Context) {
		var categories []Category
		if err := db.Find(&categories).Error; err != nil {
			c.AbortWithStatus(500)
		}
		c.JSON(200, categories)
	})

	router.POST("/api/v1/categories", func(c *gin.Context) {
		var category Category
		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Omit(clause.Associations).Create(&category)
		c.JSON(201, category)
	})

	router.GET("/api/v1/categories/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var category Category

		if err := db.Where("id = ?", id).First(&category).Error; err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(200, category)
		}
	})

	router.PUT("/api/v1/categories/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var category Category

		if err := db.Where("id = ?", id).First(&category).Error; err != nil {
			c.AbortWithStatus(404)
		} else {
			db.Save(&category)
			c.Writer.WriteHeader(http.StatusNoContent)
		}
	})

	router.DELETE("/api/v1/categories/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var category Category

		if err := db.Where("id = ?", id).First(&category).Error; err != nil {
			c.AbortWithStatus(404)
		} else {
			db.Delete(&category)
			c.Writer.WriteHeader(http.StatusNoContent)
		}
	})

	router.Run()
}
