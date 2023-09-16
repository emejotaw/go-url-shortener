package entity

import "gorm.io/gorm"

type URL struct {
	LongUrl        string
	ShortUrl       string
	NumberOfClicks int64
	Enabled        bool
	gorm.Model
}
