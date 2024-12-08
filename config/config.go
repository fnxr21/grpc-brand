package config

import (
	"log"

	"github.com/fnxr21/brand/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=root21! password=root21!Save dbname=author_mst port=4011 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Migrate schema
	if err := DB.AutoMigrate(&entities.Brand{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	
}
