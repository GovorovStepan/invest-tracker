package models

import (
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	PortfolioID uint      `json:"portfolioId" gorm:"unique"`
	Portfolio   Portfolio `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ticker      string    `json:"ticker"`
	Exchange    string    `json:"exchange"`
	Note        string    `json:"note"`
}
