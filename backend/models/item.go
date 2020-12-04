package models

import "github.com/dasdachs/inventory/database"

type Item struct {
	Id         uint     `gorm:"primaryKey autoIncrement" form:"id" json:"id"`
	Name       string   `gorm:"not null" form:"name" json:"name"`
	Quantity   uint8    `gorm:"not null" form:"quantity" json:"quantity"`
	CategoryID uint     `form:"categoryId" json:"categoryId"`
	Category   Category `form:"category" json:"category"`
}

func (i *Item) Save() error {
	if err := database.DB.Create(i).Error; err != nil {
		return err
	}
	return nil
}

func (i *Item) Update() error {
	if err := database.DB.Create(i).Error; err != nil {
		return err
	}
	return nil
}

func GetAllItems() (*[]Item, error) {
	var items []Item

	if err := database.DB.Preload("Category").Find(&items).Error; err != nil {
		return nil, err
	}

	return &items, nil
}

func GetItemById(id string) (*Item, error) {
	item := new(Item)

	if err := database.DB.Preload("Category").Where("id = ?", id).First(item).Error; err != nil {
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
