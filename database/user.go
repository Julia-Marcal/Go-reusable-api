package database

import "time"

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
