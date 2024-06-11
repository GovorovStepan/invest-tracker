package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type PortfolioRequest struct {
	Name string `json:"name" binding:"required"`
}

func GetPortfolio(context *gin.Context) {
	var portfolio models.Portfolio
	// Получение локализатора из контекста
	localizer, exists := context.Get("localizer")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Localizer not found"})
		context.Abort()
		return
	}
	localizerInstance := localizer.(*i18n.Localizer)

	id := context.Param("id")

	record := database.Instance.Where("id = ?", id).First(&portfolio)
	if record.Error != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "portfolio.get.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": record})
}

func CreatePortfolio(context *gin.Context) {
	var request PortfolioRequest
	var portfolio models.Portfolio

	// Получение локализатора из контекста
	localizer, exists := context.Get("localizer")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Localizer not found"})
		context.Abort()
		return
	}
	localizerInstance := localizer.(*i18n.Localizer)

	// Привязка данных запроса к структуре
	if err := context.ShouldBindJSON(&request); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.binding_json",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}

	userID, exists := context.Get("userID")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		context.Abort()
		return
	}

	// Присвоение значений новой записи портфолио
	portfolio.Name = request.Name
	portfolio.UserID = userID.(uint) // Предполагая, что идентификатор пользователя - это uint

	if err := database.Instance.Create(&portfolio).Error; err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "portfolio.create.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}
	message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "portfolio.create.success",
	})

	context.JSON(http.StatusCreated, gin.H{"message": message, "id": portfolio.ID})
}
func UpdatePortfolio(context *gin.Context) {
	var request PortfolioRequest
	var portfolio models.Portfolio

	// Получение локализатора из контекста
	localizer, exists := context.Get("localizer")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Localizer not found"})
		context.Abort()
		return
	}
	localizerInstance := localizer.(*i18n.Localizer)

	// Привязка данных запроса к структуре
	if err := context.ShouldBindJSON(&request); err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.binding_json",
		})
		context.JSON(http.StatusBadRequest, gin.H{"error": message})
		context.Abort()
		return
	}

	id := context.Param("id")

	record := database.Instance.Where("id  = ?", id).First(&portfolio)

	if record.Error != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "error.fetching_record",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	portfolio.Name = request.Name

	// Сохранение изменений и проверка на ошибки
	if err := database.Instance.Save(&portfolio).Error; err != nil {
		message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "portfolio.update.error",
		})
		context.JSON(http.StatusInternalServerError, gin.H{"error": message})
		context.Abort()
		return
	}

	message := localizerInstance.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "portfolio.update.success",
	})

	context.JSON(http.StatusOK, gin.H{"message": message})
}
