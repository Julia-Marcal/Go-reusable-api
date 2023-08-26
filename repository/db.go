package repository

import (
	"fmt"

	database "github.com/Julia-Marcal/reusable-api/database"
	env "github.com/Julia-Marcal/reusable-api/services/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres() *gorm.DB {
	conStr := env.SetPgConnection()
	fmt.Println("about to connect to database")
	db, err := gorm.Open(postgres.Open(conStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&database.User{})
	return db
}
