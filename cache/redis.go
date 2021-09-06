package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisAdapter struct {
	host     string
	port     string
	db       int
	password string
	ttl      time.Duration
	prefix   string
	client   *redis.Client
}

func NewRedisAdapter(host string, port string, db int, password string, ttl time.Duration, prefix string) *RedisAdapter {
	redisAdapter := &RedisAdapter{
		host:     host,
		port:     port,
		db:       db,
		password: password,
		ttl:      ttl,
		prefix:   prefix,
	}

	client := redisAdapter.getNewClient()
	redisAdapter.client = client

	return redisAdapter
}

func (adapter *RedisAdapter) getNewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     adapter.host + ":" + adapter.port,
		Password: adapter.password,
		DB:       adapter.db,
	})
}

func (adapter *RedisAdapter) getTtl() time.Duration {
	return adapter.ttl
}

func (adapter *RedisAdapter) SetTtl(ttl int) {
	adapter.ttl = time.Duration(ttl) * time.Second
}

func (adapter *RedisAdapter) getClient() *redis.Client {
	return adapter.client
}

func (adapter *RedisAdapter) ping() error {
	pong, err := adapter.client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}

func (adapter *RedisAdapter) getKey(key string) string {
	return adapter.prefix + key
}
