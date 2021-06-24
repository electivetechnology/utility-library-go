package cache

import (
	"time"

	"github.com/electivetechnology/utility-library-go/logger"
	r "github.com/go-redis/redis/v7"
)

const (
	MAX_ATTEMPTS = 5               // Default number of attempts to connect to db
	INTERVAL     = 3 * time.Second // Time between attempts
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("cache")
}

type DataCache interface {
	Set(key string, data []byte)
	SetWithTags(key string, data []byte, tags []string)
	Get(key string) []byte
}

type Cache struct {
	IsConnected bool
	Logger      logger.Logging
	MaxAttempts int
	Interval    time.Duration
	Adapter     Adapter
}

func NewCache(adapter *RedisAdapter) *Cache {
	// Setup Model
	c := Cache{
		IsConnected: false,
		Logger:      log,
		MaxAttempts: MAX_ATTEMPTS,
		Interval:    INTERVAL,
		Adapter:     adapter,
	}

	return &c
}

func (cache *Cache) Set(key string, data []byte) {
	log.Printf("Setting key %s with value %v", key, data)
	ttl := cache.Adapter.getTtl()
	cache.Adapter.getClient().Set(cache.Adapter.getKey(key), data, ttl)
}

func (cache *Cache) Get(key string) []byte {
	log.Printf("Getting key %s", key)
	val, err := cache.Adapter.getClient().Get(cache.Adapter.getKey(key)).Result()
	log.Printf("Got value from cache %v", val)

	if err != nil {
		return nil
	}

	ret := []byte(val)

	return ret
}

func (cache *Cache) SetNx(key string, data []byte) *r.BoolCmd {
	log.Printf("Trying to set NX key %s with value %v", key, data)

	ttl := cache.Adapter.getTtl()

	return cache.Adapter.getClient().SetNX(cache.Adapter.getKey(key), data, ttl)
}

func (cache *Cache) SetWithTags(key string, data []byte, tags []string) {
	log.Printf("Setting key %s with value %v and tags %s", key, data, tags)
	pipe := cache.Adapter.getClient().TxPipeline()
	ttl := cache.Adapter.getTtl()
	for _, tag := range tags {
		pipe.SAdd(tag, cache.Adapter.getKey(key))
		pipe.Expire(tag, ttl)
	}

	pipe.Set(cache.Adapter.getKey(key), data, ttl)
	_, err := pipe.Exec()
	if err != nil {
		log.Printf("Could not save cache with tags %v", err)
	}
}

func (cache *Cache) Invalidate(tags []string) {
	keys := make([]string, 0)
	for _, tag := range tags {
		k, _ := cache.Adapter.getClient().SMembers(tag).Result()
		keys = append(keys, tag)
		keys = append(keys, k...)
	}
	cache.Adapter.getClient().Del(keys...)
}

func (cache *Cache) Ping() error {
	log.Printf("Pinging cache")
	return cache.Adapter.ping()
}
