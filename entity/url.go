package entity

import "gorm.io/gorm"

type URL struct {
	LongUrl        string
	UrlKey         string
	NumberOfClicks int64
	Enabled        bool
	gorm.Model
}

func New(url, urlKey string) *URL {

	return &URL{
		LongUrl:        url,
		UrlKey:         urlKey,
		NumberOfClicks: 0,
		Enabled:        true,
	}
}
