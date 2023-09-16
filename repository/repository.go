package repository

import "github.com/emejotaw/go-url-shortener/entity"

type Repository interface {
	Create(*entity.URL) error
	FindByID(int64) (*entity.URL, error)
	FindByUrlKey(string) (*entity.URL, error)
	DeleteByID(int64) error
	Update(*entity.URL) error
}
