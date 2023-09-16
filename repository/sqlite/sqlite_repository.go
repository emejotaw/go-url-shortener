package sqlite

import "gorm.io/gorm"

type sqliteRepository struct {
	db *gorm.DB
}

func NewSqliteRepository(db *gorm.DB) *sqliteRepository {

	return &sqliteRepository{db: db}
}
