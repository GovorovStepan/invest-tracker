package models

import (
	"gorm.io/gorm"
)

type Settings struct {
	gorm.Model
	UserID   uint   `json:"userId" gorm:"unique"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Language string `json:"language"  gorm:"default:ru"`
	Currency string `json:"currency"  gorm:"default:rub"`
}
