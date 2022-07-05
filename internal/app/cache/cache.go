package cache

import (
	"sync"
	"time"
)

// Simple cache, use with handler in router.go
type Cache struct {
	sync.Mutex
	cleanupInterval time.Duration
	created         time.Time
	jsonCache       []byte
}

func NewCache() *Cache {
	return &Cache{}
}

// Write to cache
func (c *Cache) DataSet(repo []byte) {
	var wg sync.WaitGroup
	wg.Add(1)
	c.Lock()
	defer c.Unlock()

	c.cleanupInterval = 30 * time.Second
	c.created = time.Now()
	c.jsonCache = repo

	go func() {
		wg.Done()
		<-time.After(c.cleanupInterval)

		c.jsonCache = nil
	}()
	wg.Wait()
}

// Read from cache
func (c *Cache) Get() []byte {
	c.Lock()
	defer c.Unlock()
	return c.jsonCache
}
