package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type TransactionCreateRequest struct {
	Amount    uint32  `json:"amount" binding:"required"`
	Price     float32 `json:"price" binding:"required"`
	Commision float32 `json:"commision" binding:"required"`
	Type      string  `json:"type" binding:"required"`
}
type TransactionUpdateRequest struct {
	Amount    uint32  `json:"amount" binding:"required"`
	Price     float32 `json:"price" binding:"required"`
	Commision float32 `json:"commision" binding:"required"`
}

type TransactionURI struct {
	PositionID    uint `uri:"position_id"`
	TransactionID uint `uri:"transaction_id"`
}

func GetTransactions(context *gin.Context) {
	var transaction models.Transaction
	var uri TransactionURI
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

	record := database.Instance.Where("position_id = ?", uri.PositionID).Find(&transaction)
	if record.Error != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "transaction.get_all.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": transaction})

}
func CreateTransaction(context *gin.Context) {
	var request TransactionCreateRequest
	var uri TransactionURI
	var transaction models.Transaction
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

	transaction.Amount = request.Amount
	transaction.Price = request.Price
	transaction.Commision = request.Commision
	transaction.Type = request.Type
	transaction.PositionID = uri.PositionID

	if err := database.Instance.Create(&transaction).Error; err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "transaction.create.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}
	message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "transaction.create.success",
	})

	context.JSON(http.StatusCreated, gin.H{"message": message, "id": transaction.ID})
}
func UpdateTransaction(context *gin.Context) {
	var request TransactionUpdateRequest
	var uri TransactionURI
	var transaction models.Transaction
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

	record := database.Instance.Where("id  = ?", uri.TransactionID).First(&transaction)

	if record.Error != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.fetching_record",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	transaction.Amount = request.Amount
	transaction.Price = request.Price
	transaction.Commision = request.Commision

	// Сохранение изменений и проверка на ошибки
	if err := database.Instance.Save(&transaction).Error; err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "transaction.update.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "transaction.update.success",
	})

	context.JSON(http.StatusOK, gin.H{"message": message})

}
func DeleteTransaction(context *gin.Context) {
	var transaction models.Transaction
	var uri TransactionURI
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

	database.Instance.Delete(&transaction, uri.TransactionID)

	result := database.Instance.First(&transaction, uri.TransactionID)

	if result.Error == nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "transaction.delete.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
	}

	message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "transaction.delete.success",
	})
	context.JSON(http.StatusOK, gin.H{"message": message})

}
