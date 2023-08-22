package queries

import (
	database "github.com/Julia-Marcal/reusable-api/database"
	repository "github.com/Julia-Marcal/reusable-api/repository"
	_ "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(user_info *database.User) error {
	db := repository.NewSqlite()
	result := db.Create(user_info)
	return result.Error
}

func FindOne(id string) (*database.User, error) {
	db := repository.NewSqlite()
	user := &database.User{}
	result := db.First(user, "id = ?", id)
	return user, result.Error
}

func FindUsers() (int64, error) {
	db := repository.NewSqlite()
	var users []database.User
	result := db.Limit(10).Find(&users)
	return result.RowsAffected, result.Error
}

func DeleteOne(id string) (*gorm.DB){
	db := repository.NewSqlite()
	user := &database.User{}
	result := db.Delete(user, "id = ?", id)
	return result
} 