package database

import (
	"blog-post-api/graph/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "user=postgres password=admin dbname=blogest sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// Use models to create tables
func CreateTables() {
	db := Connect()
	db.AutoMigrate(model.Post{}, model.Author{}, model.Comment{}, model.Like{})
}
