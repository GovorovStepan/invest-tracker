package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type UpdateRequest struct {
	UserID   string `json:"userId"`
	Language string `json:"language"`
	Currency string `json:"currency"`
}

func GetSettings(context *gin.Context) {
	var settings models.Settings
	// Получение локализатора из контекста
	localizer, _ := context.Get("localizer")
	localizerInstance := localizer.(*i18n.Localizer)

	user_id, _ := context.Get("UserID")

	record := database.Instance.Where("user = ?", user_id).First(&settings)

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
	var request UpdateRequest
	var settings models.Settings

	// Получение локализатора из контекста
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
	record := database.Instance.Where("user = ?", request.UserID).First(&settings)

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
