package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"

	"server/controllers"
	"server/database"
	"server/middlewares"
)

func main() {
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "8080"
	}
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable "
	}
	timezone := os.Getenv("DB_TIMEZONE")
	if timezone == "" {
		timezone = "UTC"
	}
	dbHost, exists := os.LookupEnv("DB_HOST")
	if !exists {
		log.Fatal().Msg("DB_HOST is not set")
	}
	dbUser, exists := os.LookupEnv("DB_USERNAME")
	if !exists {
		log.Fatal().Msg("DB_USERNAME is not set")
	}
	dbPass, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		log.Fatal().Msg("DB_PASSWORD is not set")
	}
	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		log.Fatal().Msg("DB_NAME is not set")
	}

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass, dbName, port, sslmode, timezone)

	database.Connect(connectionString)
	database.Migrate()
	router := initRouter()
	router.Run(":8080")
}

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
	{
		token := api.Group("/token")
		{
			token.POST("/refresh", controllers.RefreshToken)
		}
		user := api.Group("/user")
		{
			user.POST("/register", controllers.RegisterUser)
			user.POST("/login", controllers.LoginUser)
		}
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/secret", controllers.Secret)
		}

		settings := api.Group("/settings").Use(middlewares.Auth())
		{
			settings.GET("/", controllers.GetSettings)
			settings.POST("/", controllers.UpdateSettings)

		}

	}
	return router
}
