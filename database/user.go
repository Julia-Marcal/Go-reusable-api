package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string     `gorm:"primaryKey"`
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

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = uuid.NewString()
	return
}
