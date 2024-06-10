package models

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	RefreshToken string `json:"refreshToken"`
	User         User   `json:"userId"`
}
