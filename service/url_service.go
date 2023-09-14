package service

import "crypto/rand"

type UrlService struct {
}

func (u *UrlService) Generate(url string) string {

	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, 5)
	rand.Read(bytes)

	for index, data := range bytes {
		bytes[index] = alphanum[data%byte(len(alphanum))]
	}

	return string(bytes)
}
