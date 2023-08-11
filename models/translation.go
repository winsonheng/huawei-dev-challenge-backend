package models

import "gorm.io/gorm"

type Translation struct {
	gorm.Model
	SourceLanguageID uint `gorm:"foreignKey; not null"`
	TargetLanguageID uint `gorm:"foreignKey; not null"`
	ClientID uint `gorm:"foreignKey; not null"`
}