package routes

import (
	"github.com/emejotaw/go-url-shortener/api"
	"github.com/emejotaw/go-url-shortener/service"
	"github.com/gin-gonic/gin"
)

func Start() {

	server := gin.Default()

	urlService := service.NewUrlService()
	urlService.Generate("www.google.com")
	controller := api.NewUrlShortenerController(urlService)

	server.GET("/short/:name", controller.Redirect)
	server.Run(":3000")
}
