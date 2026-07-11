package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]cacheEntry
	mux  sync.RWMutex
}

// Add a new entry to the cache with the given key and value.
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Unlock()
}

// reapLoop periodically removes old entries from the cache.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > interval {
				delete(c.data, key)
			}
		}
		c.mux.Unlock()
	}
}

// Get retrieves the value associated with the given key from the cache.
func (c *Cache) Get(key string) (data []byte, ok bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	entry, ok := c.data[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

// NewCache creates a new Cache with the given interval for reaping old entries.
func NewCache(interval time.Duration) *Cache {
	cache := Cache{data: map[string]cacheEntry{}, mux: sync.RWMutex{}}
	go cache.reapLoop(interval)
	return &cache
}
