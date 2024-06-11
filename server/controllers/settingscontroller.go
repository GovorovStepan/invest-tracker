package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type SettingsUpdateRequest struct {
	Language string `json:"language" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

func GetSettings(context *gin.Context) {
	var settings models.Settings
	// Получение локализатора из контекста
	localizer, exists := context.Get("localizer")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Localizer not found"})
		context.Abort()
		return
	}
	localizerInstance := localizer.(*i18n.Localizer)

	userID, exists := context.Get("userID")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		context.Abort()
		return
	}

	record := database.Instance.Where("user = ?", userID).First(&settings)

	if record.Error != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.fetching_record",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": record})
}

func UpdateSettings(context *gin.Context) {
	var request SettingsUpdateRequest
	var settings models.Settings

	// Получение локализатора из контекста
	localizer, exists := context.Get("localizer")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Localizer not found"})
		context.Abort()
		return
	}
	localizerInstance := localizer.(*i18n.Localizer)

	userID, exists := context.Get("userID")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		context.Abort()
		return
	}

	if err := context.ShouldBindJSON(&request); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.binding_json",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}
	record := database.Instance.Where("user = ?", userID).First(&settings)

	if record.Error != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.fetching_record",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	settings.Language = request.Language
	settings.Currency = request.Currency

	// Сохранение изменений и проверка на ошибки
	if err := database.Instance.Save(&settings).Error; err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.saving_record",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "settings.update.success",
	})

	context.JSON(http.StatusOK, gin.H{"message": message})

}
