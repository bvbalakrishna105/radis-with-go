package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Ping the Redis server to check if it's reachable
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// Set a key-value pair
	err = rdb.Set(context.Background(), "mykey", "Hello, Redis!", 0).Err()
	if err != nil {
		fmt.Println("Failed to set key:", err)
		return
	}
	fmt.Println("Set key successfully")

	// Get the value for a key
	val, err := rdb.Get(context.Background(), "mykey").Result()
	if err != nil {
		fmt.Println("Failed to get value for key:", err)
		return
	}
	fmt.Println("Value for key 'mykey':", val)

	// Set a key with expiration
	err = rdb.Set(context.Background(), "mykey_with_expiry", "Hello, Redis with expiry!", 10*time.Second).Err()
	if err != nil {
		fmt.Println("Failed to set key with expiry:", err)
		return
	}
	fmt.Println("Set key with expiry successfully")

	// Sleep for a while to allow the key to expire
	time.Sleep(15 * time.Second)

	// Attempt to get the value for the expired key
	val, err = rdb.Get(context.Background(), "mykey_with_expiry").Result()
	if err == redis.Nil {
		fmt.Println("Key 'mykey_with_expiry' has expired")
	} else if err != nil {
		fmt.Println("Failed to get value for key:", err)
		return
	} else {
		fmt.Println("Value for key 'mykey_with_expiry':", val)
	}

	// Close the Redis client when done
	err = rdb.Close()
	if err != nil {
		fmt.Println("Error closing Redis client:", err)
		return
	}
	fmt.Println("Closed Redis client")
}
