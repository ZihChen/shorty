package router

import (
	"github.com/gin-gonic/gin"
	"shorty/app/handler/urlhandler"
)

func LoadingRouter(r *gin.Engine) {
	urlHandler := urlhandler.NewHandler()

	api := r.Group("redirect_url")
	{
		api.POST("/", urlHandler.RedirectUrl)
	}
}
