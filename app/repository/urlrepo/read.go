package urlrepo

import (
	"go.mongodb.org/mongo-driver/bson"
	"shorty/app/model"
)

func (r *repo) FindOriginUrl(redirectUrl string) (url model.Url, err error) {
	mongodb, err := r.mongodb.GetConnection()
	if err != nil {
		return
	}

	if err = mongodb.Collection(UrlCollectionName).FindOne(*r.mongodb.GetContext(), bson.D{{
		"replaceurl", redirectUrl}}).Decode(&url); err != nil {
		return
	}

	return
}