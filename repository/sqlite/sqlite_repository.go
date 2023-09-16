package sqlite

import (
	"github.com/emejotaw/go-url-shortener/entity"
	"gorm.io/gorm"
)

type sqliteRepository struct {
	db *gorm.DB
}

func NewSqliteRepository(db *gorm.DB) *sqliteRepository {

	return &sqliteRepository{db: db}
}

func (s *sqliteRepository) Create(url *entity.URL) error {
	s.db.Create(url)
	return nil
}

func (s *sqliteRepository) Update(url *entity.URL) error {
	s.db.Save(url)
	return nil
}

func (s *sqliteRepository) FindByID(id int) (*entity.URL, error) {
	url := &entity.URL{}
	s.db.First("id = ?", id).Find(url)
	return url, nil
}

func (s *sqliteRepository) FindByUrlKey(shortUrl string) (*entity.URL, error) {
	url := &entity.URL{}
	s.db.First("urlKey = ?", shortUrl).Find(url)
	return url, nil
}

func (s *sqliteRepository) DeleteByID(id int) error {
	url := &entity.URL{}
	s.db.Where("id = ?", id).Find(url)
	s.db.Delete(url)
	return nil
}
