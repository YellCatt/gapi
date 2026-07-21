package config

import (
	"log"

	"github.com/example/gapi/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:"+cfg.Database.Path), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}