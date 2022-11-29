package urlhandler

import (
	"shorty/app/service/urlservice"
)

type Handler struct {
	urlService urlservice.Interface
}

func NewHandler() *Handler {
	return &Handler{
		urlService: urlservice.NewService(),
	}
}
