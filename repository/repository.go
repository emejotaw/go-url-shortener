package repository

import "github.com/emejotaw/go-url-shortener/entity"

type Repository interface {
	Create(*entity.URL)
	FindByID(int64)
	DeleteByID(int64)
	Update(*entity.URL)
}
