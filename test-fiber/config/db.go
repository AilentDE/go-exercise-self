package config

import (
	"fiber-clean-arch-demo/internal/domain"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	return db
}