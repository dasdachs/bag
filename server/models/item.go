package models

import "github.com/dasdachs/inventory/database"

type Item struct {
	ID         uint     `form:"id" json:"id"`
	ItemName   string   `gorm:"not null" form:"itemName" json:"ItemName"`
	Quantity   uint8    `gorm:"not null" form:"quantity" json:"quantity"`
	CategoryID uint     `form:"categoryId" json:"categoryId"`
	Category   Category `form:"category" json:"category,omitempty"`
}

func (i *Item) Save() error {
	if err := database.DB.Create(i).Error; err != nil {
		return err
	}
	return nil
}

func (i *Item) Update() error {
	if err := database.DB.Save(i).Error; err != nil {
		return err
	}
	return nil
}

func GetAllItems() (*[]Item, error) {
	var items []Item

	err := database.DB.Model(&Item{}).Select("items.id, items.item_name, items.quantity, categories.id as category_id, categories.category_name").Joins("LEFT JOIN categories on categories.id = items.category_id").Scan(&items).Error
	if err != nil {
		return nil, err
	}

	return &items, nil
}

func GetItemById(id string) (*Item, error) {
	var item Item

	if err := database.DB.Preload("Category").First(&item, id).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func GetItemByIdWithCategory(id string) (*Item, error) {
	item := new(Item)

	err := database.DB.Model(&Item{}).Select("items.id, items.item_name, items.quantity, categories.id as category_id, categories.category_name").Joins("LEFT JOIN categories on categories.id = items.category_id").Scan(&item).Error

	if err != nil {
		return nil, err
	}

	return item, nil
}

func DeleteItem(id string) error {
	item, err := GetItemById(id)
	if err != nil {
		return err
	}
	database.DB.Delete(item)
	return nil
}
