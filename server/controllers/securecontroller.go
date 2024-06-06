package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

func Secret(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "secret message"})
}
