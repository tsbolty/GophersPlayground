package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitializeRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// Verify connection
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
}

func SetUserSession(userID int, refreshToken string, accessToken string, expiration time.Duration) error {
	// Create a session key based on user ID. Example: "session:1"
	sessionKey := fmt.Sprintf("session:%d", userID)

	sessionData, marshallErr := json.Marshal(map[string]interface{}{
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
	if marshallErr != nil {
		return fmt.Errorf("failed to marshal user session data: %w", marshallErr)
	}

	// Set the session key in Redis with the refresh token as its value.
	err := Client.Set(context.Background(), sessionKey, sessionData, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set user session in Redis: %w", err)
	}

	fmt.Printf("User session set in Redis for userID %d\n", userID)
	return nil
}

func DeleteUserSession(userID int) error {
	sessionKey := fmt.Sprintf("session:%d", userID)

	_, err := Client.Del(context.Background(), sessionKey).Result()
	if err != nil {
		return err
	}

	return nil
}
