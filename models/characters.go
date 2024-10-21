package models

import "gorm.io/gorm"

type Characters struct {
	ID          uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	ImageUrl    *string `json:"image_url"`
	Tag         *string `json:"tag"`
}

func MigrateCharacters(db *gorm.DB) error {
	err := db.AutoMigrate(&Characters{})
	return err
}
