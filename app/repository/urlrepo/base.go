package urlrepo

import (
	"shorty/app/model"
	"shorty/internal/mongodb"
	"sync"
)

type Interface interface {
	InsertData(url model.Url) (err error)
	FindOriginUrl(redirectUrl string) (url model.Url, err error)
}

type repo struct {
	mongodb mongodb.Interface
}

const UrlCollectionName = "url"

var singleton *repo
var once sync.Once

func NewRepo() Interface {
	once.Do(func() {
		singleton = &repo{
			mongodb: mongodb.NewInstance(),
		}
	})
	return singleton
}
