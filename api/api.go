package api

import (
	"fmt"
	"net/http"

	"github.com/emejotaw/go-url-shortener/service"
	"github.com/gin-gonic/gin"
)

type urlShortenerController struct {
	urlService *service.UrlService
}

func NewUrlShortenerController(urlService *service.UrlService) *urlShortenerController {
	return &urlShortenerController{urlService: urlService}
}

func (u *urlShortenerController) Redirect(c *gin.Context) {

	name := c.Param("name")

	fmt.Println(name)

	url, err := u.urlService.Get(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": err.Error(),
		})
		return
	}
	http.RedirectHandler(url, http.StatusMovedPermanently)
}
