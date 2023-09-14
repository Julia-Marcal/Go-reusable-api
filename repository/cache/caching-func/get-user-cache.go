package cachingfunc

import (
	"context"
	"fmt"
	"time"

	"github.com/Julia-Marcal/reusable-api/database"
	cache "github.com/Julia-Marcal/reusable-api/repository/cache"
)

func GetCachedUser(email string) (database.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	RedisClient, _ := cache.RedisInit()

	userData, err := RedisClient.HGetAll(ctx, "user:"+email).Result()
	if err != nil {
		return database.User{}, err
	}

	if len(userData) == 0 {
		return database.User{}, fmt.Errorf("user not found in cache")
	}

	var user database.User

	return user, nil
}
