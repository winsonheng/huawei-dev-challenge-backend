package models

import "gorm.io/gorm"

type Text struct {
	gorm.Model
	Content string `gorm:"not null; unique"`
	LanguageID uint `gorm:"foreignKey; not null"`
}