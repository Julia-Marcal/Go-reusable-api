package repository

import (
	"fmt"

	database "github.com/Julia-Marcal/reusable-api/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres() *gorm.DB {
	connectionStr := "user=postgres password=password dbname=api_db host=postgres port=5432 sslmode=disable"
	fmt.Println("about to connect to database")
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&database.User{})
	return db
}
