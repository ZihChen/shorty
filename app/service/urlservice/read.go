package urlservice

func (s *service) GetOriginUrl(redirectUrl string) (originUrl string, err error) {
	url, err := s.urlRepo.FindOriginUrl(redirectUrl)
	if err != nil {
		return
	}

	return url.OriginUrl, nil
}
