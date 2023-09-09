package cachingfunc

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Julia-Marcal/reusable-api/database"
	cache "github.com/Julia-Marcal/reusable-api/repository/cache"
)

func CacheUser(user database.User) error {

	cacheExpiration := 7 * 24 * time.Hour // 7 days

	ctx := context.Background()

	RedisClient := cache.RedisInit()
	userMap := map[string]interface{}{
		"id":        user.Id,
		"name":      user.Name,
		"lastName":  user.LastName,
		"age":       user.Age,
		"email":     user.Email,
		"password":  user.Password,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
	}

	userDataJSON, err := json.Marshal(userMap)
	if err != nil {
		return err
	}

	redisKey := "user:" + user.Email
	err = RedisClient.Set(ctx, redisKey, userDataJSON, cacheExpiration).Err()
	if err != nil {
		return err
	}

	return nil

}