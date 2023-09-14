package main

import (
	"fmt"

	"github.com/emejotaw/go-url-shortener/service"
)

func main() {

	url := "https://google.com"

	urlService := service.UrlService{}
	shortURL := urlService.Generate(url)

	fmt.Println(shortURL)
}
