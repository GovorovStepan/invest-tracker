package models

import (
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	PortfolioID uint      `json:"portfolioId"`
	Portfolio   Portfolio `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ticker      string    `json:"ticker"`
	Exchange    string    `json:"exchange"`
	Note        string    `json:"note"`
}
