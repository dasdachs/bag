package models

type Category struct {
	Id    uint   `gorm:"primaryKey autoIncrement" form:"id" json:"id"`
	Name  string `gorm:"not null unique" form:"name" json:"name"`
	Items []Item `form:"category" json:"items"`
}
