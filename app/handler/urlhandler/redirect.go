package urlhandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) RedirectUrl(ctx *gin.Context) {
	redirectUrl := ctx.Param("redirect_url")
	if redirectUrl == "" {
		ctx.JSON(http.StatusOK, "redirect url required")
		return
	}

	originUrl, err := h.urlService.GetOriginUrl(redirectUrl)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, originUrl)
	ctx.Abort()
}
