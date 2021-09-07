package cache

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
)

type Adapter interface {
	getTtl() time.Duration
	SetTtl(int)
	getClient() *redis.Client
	ping() error
	getKey(key string) string
}

func NewCacheAdapter() *RedisAdapter {
	host, port, db, password, ttl, prefix := getAdapterDetails()
	a := NewRedisAdapter(host, port, db, password, ttl, prefix)

	return a
}

func getAdapterDetails() (host string, port string, db int, password string, ttl time.Duration, prefix string) {
	// Get Cache env
	host = os.Getenv("CACHE_HOST")
	if host == "" {
		host = "redis"
	}

	port = os.Getenv("CACHE_PORT")
	if port == "" {
		port = "6379"
	}

	cdb := os.Getenv("CACHE_DB")
	if cdb == "" {
		cdb = "0"
	}

	password = os.Getenv("CACHE_PASSWORD")
	if password == "" {
		password = ""
	}

	cttl := os.Getenv("CACHE_TTL")
	if cttl == "" {
		cttl = "3600"
	}

	prefix = os.Getenv("CACHE_PREFIX")
	if prefix == "" {
		prefix = ""
	}

	db, _ = strconv.Atoi(cdb)
	ttlInt, _ := strconv.Atoi(cttl)

	return host, port, db, password, time.Duration(ttlInt) * time.Second, prefix
}
