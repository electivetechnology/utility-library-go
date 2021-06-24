package cache

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type Adapter interface {
	getTtl() time.Duration
	getClient() *redis.Client
	ping() error
	getKey(key string) string
}
