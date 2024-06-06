package controllers

import (
	"auth/database"
	"auth/models"
	"auth/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func RefreshToken(context *gin.Context) {
	var request TokenRefreshRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	claims, err := token.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		context.JSON(401, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Where("id = ?", claims.UserID).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	accessTokenString, err := token.GenerateAccessToken(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"accessToken": accessTokenString})

}
