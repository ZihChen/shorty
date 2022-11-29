package urlservice

import (
	"shorty/app/global/helper"
	"shorty/app/model"
)

func (s *service) CreateRedirectUrl(originUrl string) (redirectUrl string, err error) {
	replaceUrl, err := helper.GenerateReplaceUrl(originUrl)
	if err != nil {
		return
	}

	if err = s.urlRepo.InsertData(model.Url{
		OriginUrl:  originUrl,
		ReplaceUrl: replaceUrl}); err != nil {
		return
	}

	return
}
