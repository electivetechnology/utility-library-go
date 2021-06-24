package cache

import (
	"testing"
	"time"
)

func TestNewRedisAdapter(t *testing.T) {
	host := "localhost"
	port := "6379"
	db := 0
	ttl := 5 * time.Second
	adapter := NewRedisAdapter(host, port, db, "", ttl, "")

	if adapter.host != host {
		t.Errorf("NewRedisAdapter() failed, expected host to be %s, got %s", host, adapter.host)
	}

	if adapter.db != db {
		t.Errorf("NewRedisAdapter() failed, expected db to be %d, got %d", db, adapter.db)
	}

	if adapter.ttl != ttl {
		t.Errorf("NewRedisAdapter() failed, expected host to be %d, got %d", int64(ttl/time.Second), int64(adapter.ttl/time.Second))
	}
}
