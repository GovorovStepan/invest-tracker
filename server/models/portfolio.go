package models

import (
	"gorm.io/gorm"
)

type Portfolio struct {
	gorm.Model
	UserID uint   `json:"userId" gorm:"unique"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name   string `json:"name"`
}
