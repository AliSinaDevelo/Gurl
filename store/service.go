package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	storeService.redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Gurl is feeling lonely, can't connect to Redis: %v", err))
	}

	fmt.Printf("\nGurl is happy, connected to Redis: %v\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveURLMapping(shortUrl string, originalUrl string, userID string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Gurl is feeling lonely, can't save URL mapping: %v", err))
	}
}

func RetrieveInitialURL(shortUrl string) string {
	val, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Gurl is feeling lonely, can't retrieve URL mapping: %v", err))
	}
	return val
}
