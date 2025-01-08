package db

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(Ctx).Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Successfully, connected to redis!")
}
