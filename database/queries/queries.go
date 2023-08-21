package queries

import (
	database "github.com/Julia-Marcal/reusable-api/database"
	repository "github.com/Julia-Marcal/reusable-api/repository"
	_ "github.com/gin-gonic/gin"
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
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
