package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

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
		log.Fatal("DB_HOST is not set")
	}
	dbUser, exists := os.LookupEnv("DB_USER")
	if !exists {
		log.Fatal("DB_USER is not set")
	}
	dbPass, exists := os.LookupEnv("DB_PASS")
	if !exists {
		log.Fatal("DB_PASS is not set")
	}
	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		log.Fatal("DB_NAME is not set")
	}

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass, dbName, port, sslmode, timezone)

	database.Connect(connectionString)
	database.Migrate()
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", controllers.Health)

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

	}
	return router
}
