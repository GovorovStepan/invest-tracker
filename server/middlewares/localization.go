package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Localization(bundle *i18n.Bundle) gin.HandlerFunc {
	return func(context *gin.Context) {
		lang := context.GetHeader("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang)
		context.Set("localizer", localizer)
		context.Next()
	}
}
