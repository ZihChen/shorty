package urlrepo

import "shorty/app/model"

func (r *repo) InsertData(url model.Url) (err error) {
	mongodb, err := r.mongodb.GetConnection()
	if err != nil {
		return
	}

	_, err = mongodb.Collection(UrlCollectionName).InsertOne(*r.mongodb.GetContext(), url)
	if err != nil {
		return
	}

	return
}
