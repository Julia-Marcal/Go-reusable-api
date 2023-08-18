package repository

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        string     `gorm:"primaryKey;size:50"`
	Name      string     `gorm:"not null;size:50" sql:"index"`
	LastName  string     `gorm:"not null;size:50"`
	Age       int32      `gorm:"not null"`
	Email     string     `gorm:"not null"`
	CreatedAt *time.Time `gorm:"default:current_timestamp"`
	UpdatedAt *time.Time `gorm:"default:current_timestamp"`
}

func (User) TableName() string {
	return "users"
}

func NewPostgres() *gorm.DB {
	fmt.Println("about to connect to database")
	fmt.Printf("%v", os.Getenv("database"))
	db, err := gorm.Open(postgres.Open(os.Getenv("database")), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&User{})
	return db
}
