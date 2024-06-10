package models

import (
	"gorm.io/gorm"
)

type Settings struct {
	gorm.Model
	User     User   `json:"userId" gorm:"unique"`
	Language string `json:"language"  gorm:"default:ru"`
	Currency string `json:"currency"  gorm:"default:rub"`
}
