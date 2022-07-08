package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func main() {
	var redisHost = "localhost:6379"
	var redisPassword = ""

	ctx := context.Background()

	rdb := newRedisClient(redisHost, redisPassword)
	fmt.Println("redis client initialized")

	ttl := time.Duration(60) * time.Second

	// Set Value Key to Redis
	op1 := rdb.Set(ctx, "newKey", "NewData", ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}

	// Get Value Key Redis
	val := rdb.Get(ctx, "newKey")
	fmt.Println(val)
}
