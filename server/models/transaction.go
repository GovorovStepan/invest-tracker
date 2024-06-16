package models

type Transaction struct {
	Model
	PositionID uint    `json:"positionId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount     uint32  `json:"amount"`
	Price      float32 `json:"price"`
	Commision  float32 `json:"commision"`
	Type       string  `json:"type"`
}
