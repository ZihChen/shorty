package urlhandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) RedirectUrl(ctx *gin.Context) {
	url := ctx.PostForm("url")

	redirectUrl, err := h.urlService.CreateRedirectUrl(url)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, redirectUrl)
}
