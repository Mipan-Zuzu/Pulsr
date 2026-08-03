package cache

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"os"

	"github.com/redis/go-redis/v9"
)

var (
	instance *redis.Client
	once sync.Once
)

func GetClient() *redis.Client {
	once.Do(func ()  {
		url := os.Getenv("UPSTASH_REDIS_URL")

		opt, err := redis.ParseURL(url)
		if err != nil {
			log.Fatalf("invalid redis url: %s", err)
		}
		instance = redis.NewClient(opt)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) 
		defer cancel()

		if err := instance.Ping(ctx).Err(); err != nil {
			log.Fatalf("invalid connect to redis %s", err)
		}
		fmt.Printf("Succses connect to redis")
	})
	fmt.Printf("Succses connect to redis")
	return instance
}