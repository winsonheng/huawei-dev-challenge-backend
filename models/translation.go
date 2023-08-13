package models

import "gorm.io/gorm"

type Translation struct {
	gorm.Model
	SourceTextID uint `gorm:"foreignKey; not null"`
	TargetTextID uint `gorm:"foreignKey; not null"`
	ClientID uint `gorm:"foreignKey; not null"`
}