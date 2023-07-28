package shorter

import (
	"crypto/rand"
	"math/big"
)

type shortedLinks map[string]string

type Shorter struct {
	links shortedLinks
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
	s.links[shortLink] = link
	return shortLink
}

func (s *Shorter) GetFullUrl(shortString string) (string, bool) {
	value, exist := s.links[shortString]
	if exist {
		return value, exist
	}
	return "", exist
}

func getShortLink() string {
	shortedUrl, err := generateRandomString(10)
	if err != nil {
		panic(err)
	}
	return shortedUrl
}

func GetShorter() *Shorter {
	m := make(map[string]string)
	return &Shorter{links: m}
}
