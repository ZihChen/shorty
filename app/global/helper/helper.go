package helper

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

func GenerateReplaceUrl(originUrl string) (replaceUrl string, err error) {
	var letters = []rune(originUrl)
	s := make([]rune, 8)
	for i := range s {
		var key *big.Int
		key, err = rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return
		}
		s[i] = letters[key.Int64()]
	}
	return base64.StdEncoding.EncodeToString([]byte(string(s))), nil
}
