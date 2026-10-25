package pokecache

import (
	//"fmt"
	"time"
	"sync"
)

type Cache struct {
	mu sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	r, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return r.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	

	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.entries {
			time := time.Since(value.createdAt)
			if time > interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}