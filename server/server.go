package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/dasdachs/inventory/controllers"
	"github.com/dasdachs/inventory/database"
	"github.com/dasdachs/inventory/models"
)

//go:embed _public/*
var static embed.FS

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	db.AutoMigrate(&models.Category{}, &models.Item{})

	defer db.Close()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		itemsGroup := v1.Group("/items")
		{
			itemsGroup.GET("", controllers.GetAllItems)
			itemsGroup.POST("", controllers.CreateItem)
			itemsGroup.GET("/:id", controllers.GetItemById)
			itemsGroup.PUT("/:id", controllers.UpdateItem)
			itemsGroup.DELETE("/:id", controllers.DeleteItem)
		}
		categoriesGroup := v1.Group("/categories")
		{
			categoriesGroup.GET("", controllers.GetAllCategories)
			categoriesGroup.POST("", controllers.CreateCategory)
			categoriesGroup.GET("/:categoryId", controllers.GetCategoryById)
			categoriesGroup.PUT("/:categoryId", controllers.UpdateCategory)
			categoriesGroup.DELETE("/:categoryId", controllers.DeleteCategory)
			categoriesGroup.GET("/:categoryId/items", controllers.GetItemById)
			categoriesGroup.POST("/:categoryId/items", controllers.AppendItemToCategory)
			categoriesGroup.PUT("/:categoryId/items/:id", controllers.UpdateItem)
			categoriesGroup.DELETE("/:categoryId/items/:id", controllers.DeleteItem)
		}
	}

	router.StaticFS("/", mustFS())

	if err := router.Run(); err != nil {
		panic("Failed to start server")
	}
}

func mustFS() http.FileSystem {
	sub, err := fs.Sub(static, "_public")

	if err != nil {
		panic(err)
	}

	return http.FS(sub)
}
