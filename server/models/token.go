package models

type Token struct {
	Model
	RefreshToken string `json:"refreshToken"`
	User         User   `json:"userId"`
}
