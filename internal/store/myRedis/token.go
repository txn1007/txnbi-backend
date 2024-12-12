package myRedis

import (
	"context"
	"fmt"
	"time"
)

func SetUserToken(userID int64, token string) {
	key := fmt.Sprintf("user-token-current-%d", userID)
	Cli.Set(context.Background(), key, token, 7*24*time.Hour)
}

func GetUserToken(userID int64) (string, error) {
	key := fmt.Sprintf("user-token-current-%d", userID)
	return Cli.Get(context.Background(), key).Result()
}

func DeleteUserToken(userID int64) {
	key := fmt.Sprintf("user-token-current-%d", userID)
	Cli.Del(context.Background(), key)
}
