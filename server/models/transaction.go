package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	PositionID uint     `json:"positionId"`
	Position   Position `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount     uint32   `json:"amount"`
	Price      float32  `json:"price"`
	Commision  float32  `json:"commision"`
	Type       string   `json:"type"`
}
