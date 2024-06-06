package main

import (
	"github.com/gin-gonic/gin"

	"auth/controllers"
	"auth/database"
	"auth/middlewares"
)

func main() {
	database.Connect("host=localhost user=root password=root dbname=gsp port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh")
	database.Migrate()
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
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
