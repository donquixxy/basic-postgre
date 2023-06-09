package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:32769",
		Password: "",
		DB:       0,
	})

	er := client.Ping(context.TODO())

	if er != nil {
		fmt.Println(er)
	}

	return client
}
