package service

import (
	"crypto/rand"
	"errors"
	"log"
)

type UrlService struct {
	shortUrls map[string]string
}

func NewUrlService() *UrlService {
	return &UrlService{
		shortUrls: make(map[string]string),
	}
}

func (u *UrlService) Get(dns string) (string, error) {

	if url, ok := u.shortUrls[dns]; ok {
		return url, nil
	}

	return "", errors.New("url not found")
}

func (u *UrlService) Generate(url string) string {

	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, 5)
	rand.Read(bytes)

	for index, data := range bytes {
		bytes[index] = alphanum[data%byte(len(alphanum))]
	}

	shortName := string(bytes)
	log.Println("Generated name: ", shortName)

	u.shortUrls[shortName] = url
	return string(bytes)
}
