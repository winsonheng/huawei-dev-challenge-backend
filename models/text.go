package models

import "gorm.io/gorm"

type Text struct {
	gorm.Model
	Content string `gorm:"not null"`
	LanguageID uint `gorm:"foreignKey; not null"`
}