package main

import (
	"fmt"
	"log"

	"github.com/dasdachs/inventory/database"
	"github.com/dasdachs/inventory/models"
)

func test() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	db.AutoMigrate(&models.Category{}, &models.Item{})

	defer db.Close()

	item := models.Item{
		ItemName: "Mleko",
		Quantity: 12,
	}

	category := models.Category{
		CategoryName: "Mlečni izdelki",
	}

	fmt.Printf("%v", item)
	fmt.Printf("%v", category)

	db.Create(&category)
	db.Create(&item)

	fmt.Printf("%v", item)
	fmt.Printf("%v", category)

	// Start association mode
	if err := database.DB.Model(&category).Association("Items").Error; err != nil {
		log.Fatal(err)
	}

	db.Model(&category).Association("Items").Append(&item)

	fmt.Printf("%v", item)
	fmt.Printf("%v", category)
}
