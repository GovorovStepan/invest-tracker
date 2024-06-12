package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type PositionRequest struct {
	Ticker   string `json:"ticker" binding:"required"`
	Exchange string `json:"exchange" binding:"required"`
	Note     string `json:"note"`
}

type PositionURI struct {
	PortfolioID uint `uri:"portfolio_id" binding:"uint"`
	PositionID  uint `uri:"position_id" binding:"uint"`
}

func GetPositionByID(context *gin.Context) {
	var position models.Position
	var uri PositionURI
	localizer, exists := context.Get("localizer")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Localizer not found"})
		context.Abort()
		return
	}
	localizerInstance := localizer.(*i18n.Localizer)

	// Привязка данных запроса к структуре
	if err := context.ShouldBindUri(&uri); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.invalid_id",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}

	record := database.Instance.Where("id = ?", uri.PositionID).First(&position)
	if record.Error != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "position.get_by_id.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": record})

}
func GetPositions(context *gin.Context) {
	var position models.Position
	var uri PositionURI
	localizer, exists := context.Get("localizer")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Localizer not found"})
		context.Abort()
		return
	}
	localizerInstance := localizer.(*i18n.Localizer)

	// Привязка данных запроса к структуре
	if err := context.ShouldBindUri(&uri); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.invalid_id",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}

	records := database.Instance.Where("portfolio_id = ?", uri.PortfolioID).Find(&position)
	if records.Error != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "position.get_all.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": records})

}
func CreatePosition(context *gin.Context) {
	var request PositionRequest
	var uri PositionURI
	var position models.Position
	localizer, exists := context.Get("localizer")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Localizer not found"})
		context.Abort()
		return
	}
	localizerInstance := localizer.(*i18n.Localizer)

	// Привязка данных запроса к структуре
	if err := context.ShouldBindUri(&uri); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.invalid_id",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}

	// Привязка данных запроса к структуре
	if err := context.ShouldBindJSON(&request); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.binding_json",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}

	position.Ticker = request.Ticker
	position.Exchange = request.Exchange
	position.Note = request.Note
	position.PortfolioID = uri.PortfolioID

	if err := database.Instance.Create(&position).Error; err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "position.create.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}
	message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "position.create.success",
	})

	context.JSON(http.StatusCreated, gin.H{"message": message, "id": position.ID})
}
func UpdatePosition(context *gin.Context) {
	var request PositionRequest
	var uri PositionURI
	var position models.Position
	localizer, exists := context.Get("localizer")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Localizer not found"})
		context.Abort()
		return
	}
	localizerInstance := localizer.(*i18n.Localizer)

	// Привязка данных запроса к структуре
	if err := context.ShouldBindUri(&uri); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.invalid_id",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}

	// Привязка данных запроса к структуре
	if err := context.ShouldBindJSON(&request); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.binding_json",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}

	record := database.Instance.Where("id  = ?", uri.PositionID).First(&position)

	if record.Error != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.fetching_record",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	position.Ticker = request.Ticker
	position.Exchange = request.Exchange
	position.Note = request.Note

	// Сохранение изменений и проверка на ошибки
	if err := database.Instance.Save(&position).Error; err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "position.update.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "position.update.success",
	})

	context.JSON(http.StatusOK, gin.H{"message": message})

}
func DeletePosition(context *gin.Context) {
	var position models.Position
	var uri PositionURI
	localizer, exists := context.Get("localizer")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Localizer not found"})
		context.Abort()
		return
	}
	localizerInstance := localizer.(*i18n.Localizer)

	// Привязка данных запроса к структуре
	if err := context.ShouldBindUri(&uri); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.invalid_id",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}

	database.Instance.Delete(&position, uri.PositionID)

	result := database.Instance.First(&position, uri.PositionID)

	if result.Error == nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "position.delete.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
	}

	message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "position.delete.success",
	})
	context.JSON(http.StatusOK, gin.H{"message": message})

}
