package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	
	// if err := db.AutoMigrate(&User{}); err != nil {
	// 	log.Fatalf("auto-migrate failed: %v", err)
	// }

	return db
}
