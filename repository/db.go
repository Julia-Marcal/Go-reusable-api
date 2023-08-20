package repository

import (
	"fmt"

	database "github.com/Julia-Marcal/reusable-api/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite() *gorm.DB {
	fmt.Println("about to connect to database")
	db, err := gorm.Open(sqlite.Open("golang-db.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&database.User{})
	return db
}
