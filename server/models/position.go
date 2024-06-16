package models

type Position struct {
	Model
	PortfolioID uint   `json:"portfolioId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ticker      string `json:"ticker"`
	Exchange    string `json:"exchange"`
	Note        string `json:"note"`
}
