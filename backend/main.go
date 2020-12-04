package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/dasdachs/inventory/controllers"
	"github.com/dasdachs/inventory/database"
	"github.com/dasdachs/inventory/models"
)

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
			itemsGroup.GET("/", controllers.GetAllItems)
			itemsGroup.POST("/", controllers.CreateOrUpdateItem)
			itemsGroup.GET("/:id", controllers.GetItemById)
			itemsGroup.PUT("/:id", controllers.CreateOrUpdateItem)
			itemsGroup.DELETE("/:id", controllers.DeleteItem)
		}
	}
	// 	v1.GET("/", controllers.GetAllItems)
	// v1.POST("/",)
	// v1.GET("/:id",)
	// v1.PUT("/:id",)
	// v1.DELETE(":id",)
	// }
	// categoriesGroup := v1.Group("/categories") {
	// 	v1.GET("/", GetAllItems)
	// 	v1.POST("/",)
	// 	v1.GET("/:id",)
	// 	v1.PUT("/:id",)
	// 	v1.DELETE(":id",)
	// }
	router.Run()
}
