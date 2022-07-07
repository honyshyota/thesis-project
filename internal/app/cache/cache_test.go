package cache_test

import (
	"main/internal/app/cache"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache_DataSet(t *testing.T) {
	cache := cache.TestNullCache(t)
	cacheObj := []byte("some cache")
	cache.DataSet(cacheObj)
	assert.NotEmpty(t, cache)
}

func TestCache_Get(t *testing.T) {
	cache := cache.TestCache(t)
	obj := cache.Get()
	assert.NotEmpty(t, obj)
}
