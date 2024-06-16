package models

type Portfolio struct {
	Model
	UserID uint   `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name   string `json:"name"`
}
