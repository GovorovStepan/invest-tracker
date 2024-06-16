package models

type Settings struct {
	Model
	UserID   uint   `json:"userId" gorm:"unique; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Language string `json:"language"  gorm:"default:ru"`
	Currency string `json:"currency"  gorm:"default:rub"`
}
