package models

import "github.com/dasdachs/inventory/database"

type Category struct {
	ID           uint   `form:"id" json:"id,omitempty"`
	CategoryName string `gorm:"unique" form:"categoryName" json:"categoryName,omitempty"`
	Items        []Item `form:"category" json:"items,omitempty"`
}

func (c *Category) SaveCategory() error {
	if err := database.DB.Create(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) UpdateCategory() error {
	if err := database.DB.Save(c).Error; err != nil {
		return err
	}
	return nil
}

func AppendItem(id string, item *Item) error {
	category, err := GetCategoryById(id)

	if err != nil {
		return err
	}
	if err := database.DB.Model(&category).Association("Items").Error; err != nil {
		return err
	}

	database.DB.Model(&category).Association("Items").Append(item)

	return nil
}

func AppendItemById(id string, itemId string) error {
	category, err := GetCategoryById(id)
	item, itemErr := GetItemById(itemId)

	if err != nil {
		return err
	}

	if itemErr != nil {
		return itemErr
	}
	if err := database.DB.Model(&category).Association("Items").Error; err != nil {
		return err
	}

	database.DB.Model(&category).Association("Items").Append(item)

	return nil
}

func GetAllCategories() (*[]Category, error) {
	var categories []Category

	if err := database.DB.Preload("Items").Find(&categories).Error; err != nil {
		return nil, err
	}

	return &categories, nil
}

func GetCategoryById(id string) (*Category, error) {
	category := new(Category)

	if err := database.DB.Preload("Items").Where("id = ?", id).First(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func DeleteCategory(id string) error {
	category, err := GetCategoryById(id)
	if err != nil {
		return err
	}
	database.DB.Delete(category)
	return nil
}
