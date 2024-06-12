package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func addTokenRoutes(rg *gin.RouterGroup) {

	token := rg.Group("/token")
	{
		token.POST("/refresh", controllers.RefreshToken)
	}

}
