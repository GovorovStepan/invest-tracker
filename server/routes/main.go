package routes

import (
	"encoding/json"
	"server/controllers"
	"server/middlewares"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Run will start the server
func Run() {
	router := initRouter()
	router.Run(":8080")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func initRouter() *gin.Engine {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("i18n/en.json")
	bundle.MustLoadMessageFile("i18n/ru.json")

	router := gin.Default()
	router.Use(logger.SetLogger())
	router.Use(middlewares.Localization(bundle))

	router.GET("/health", controllers.Health)
	router.GET("/test", controllers.GetSettings)

	api := router.Group("/api")

	addTokenRoutes(api)
	addUserRoutes(api)
	addSettingsRoutes(api)
	addPortfolioRoutes(api)

	return router
}
