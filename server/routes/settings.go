package routes

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

func addSettingsRoutes(rg *gin.RouterGroup) {

	settings := rg.Group("/settings").Use(middlewares.Auth())
	{
		settings.GET("/", controllers.GetSettings)
		settings.PUT("/", controllers.UpdateSettings)

	}
}
