package redis

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type Client interface {
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
