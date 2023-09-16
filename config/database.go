package config

import (
	"log"

	"github.com/emejotaw/go-url-shortener/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {

	dsn := "database.db"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Could not establish database connection, error: %v", err.Error())
		panic(err)
	}

	return db
}

func AutoMigrate() {

	db := GetConnection()

	db.AutoMigrate(&entity.URL{})
}
