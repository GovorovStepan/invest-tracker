package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"

	"server/database"
	"server/models"
	"server/token"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(context *gin.Context) {
	localizer, _ := context.Get("localizer")
	localizerInstance := localizer.(*i18n.Localizer)
	var user models.User
	var settings models.Settings

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	err := database.Instance.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		settings.UserID = user.ID
		if err := tx.Create(&settings).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "user.create.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})

		context.Abort()
		return
	}

	accessTokenString, err := token.GenerateAccessToken(user.Email, fmt.Sprint(user.ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	refreshTokenString, err := token.GenerateRefreshToken(fmt.Sprint(user.ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "user.create.success",
	})

	context.JSON(http.StatusCreated, gin.H{"message": message, "accessToken": accessTokenString, "refreshToken": refreshTokenString})
}

func LoginUser(context *gin.Context) {
	var request LoginRequest
	var user models.User
	localizer, _ := context.Get("localizer")
	localizerInstance := localizer.(*i18n.Localizer)
	if err := context.ShouldBindJSON(&request); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.binding_json",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}
	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "user.login.error.email",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "user.login.error.password",
		})
		context.JSON(http.StatusUnauthorized, gin.H{"error": message})
		context.Abort()
		return
	}
	accessTokenString, err := token.GenerateAccessToken(user.Email, fmt.Sprint(user.ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	refreshTokenString, err := token.GenerateRefreshToken(fmt.Sprint(user.ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"accessToken": accessTokenString, "refreshToken": refreshTokenString})
}
