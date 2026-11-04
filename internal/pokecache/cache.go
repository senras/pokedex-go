package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    *sync.RWMutex
	items map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mu:    &sync.RWMutex{},
		items: make(map[string]CacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.items {
			if time.Since(entry.createdAt) > interval {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = CacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.items[key]
	return val.value, ok
}
