package urlservice

import (
	"shorty/app/repository/urlrepo"
	"sync"
)

type Interface interface {
	CreateRedirectUrl(originUrl string) (redirectUrl string, err error)
}

type service struct {
	urlRepo urlrepo.Interface
}

var singleton *service
var once sync.Once

func NewService() Interface {
	once.Do(func() {
		singleton = &service{
			urlRepo: urlrepo.NewRepo(),
		}
	})
	return singleton
}
