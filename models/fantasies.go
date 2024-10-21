package models

import "gorm.io/gorm"

type Fantasies struct {
	ID           uint    `gorm:"primary key;autoIncrement" json:"id"`
	Title        *string `json:"title"`
	Description  *string `json:"description"`
	ImageUrl     *string `json:"image_url"`
	Likes        *int    `json:"likes"`
	MessageCount *int    `json:"message_count"`
	Tag          *string `json:"tag"`
}

func MigrateFantasies(db *gorm.DB) error {
	err := db.AutoMigrate(&Characters{})
	return err
}
