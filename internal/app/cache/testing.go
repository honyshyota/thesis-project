package cache

import (
	"testing"
	"time"
)

func TestNullCache(t *testing.T) *Cache {
	return &Cache{}
}

func TestCache(t *testing.T) *Cache {
	return &Cache{
		cleanupInterval: time.Hour,
		created:         time.Now(),
		jsonCache:       []byte("some cache"),
	}
}
