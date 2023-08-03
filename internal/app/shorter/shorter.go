package shorter

import (
	"crypto/rand"
	"math/big"
)

type shortedLinks map[string]string

type Shorter struct {
	Links shortedLinks
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func (s *Shorter) AddLink(link string) string {
	shortLink := getShortLink()
	s.Links[shortLink] = link
	return shortLink
}

func (s *Shorter) GetFullURL(shortString string) (string, bool) {
	value, exist := s.Links[shortString]
	if exist {
		return value, exist
	}
	return "", exist
}

func getShortLink() string {
	shortedURL, err := generateRandomString(10)
	if err != nil {
		panic(err)
	}
	return shortedURL
}

func GetShorter() *Shorter {
	m := make(map[string]string)
	return &Shorter{Links: m}
}
