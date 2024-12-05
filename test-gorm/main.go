package main

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MixinTime struct {
  CreatedAt time.Time `gorm:"autoCreateTime"`
  UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Product struct {
	ID   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Code  string `gorm:"size:128;uniqueIndex"`
	Price uint `gorm:"index"`
  MixinTime
}

func main() {
  time.Local = time.UTC
	dsn := "host=localhost user=root password=rootpassword dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
  })
	if err != nil {
		panic("failed to connect database")
	}
  // Connection pool settings
  // sqlDB, err := db.DB()
  // sqlDB.SetMaxIdleConns(10)
  // sqlDB.SetMaxOpenConns(100)
  // sqlDB.SetConnMaxLifetime(time.Hour)

  // uuid4 extension
  err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
  if err != nil {
    panic("failed to create extension")
  }

	// Recreate schema
	db.Migrator().DropTable(&Product{}) // Drop the table if it exists
	err = db.AutoMigrate(&Product{}) // Recreate the table
  if err != nil {
    panic("failed to migrate")
  }

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  // db.First(&product, 1) // by primary key
  db.First(&product, "code = ?", "D42") // where condition

	// Update
	db.Model(&product).Update("Price", 200)
	// Update multiple fields
	db.Model(&product).Updates(Product{Code: "F42", Price: 300})
	db.Model(&product).Updates(map[string]interface{}{"Code": "D42", "Price": 400, })

	// Delete
	// db.Delete(&product, "code = ?", "D42")

  // Close the connection
  // sqlDB.Close()
}
