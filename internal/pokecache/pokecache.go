package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]CacheEntry
	mux   sync.Mutex
}

// New -
func New(interval time.Duration) *Cache {
	cache := &Cache{
		cache: make(map[string]CacheEntry),
		mux:   sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

// Get -
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if item, ok := c.cache[key]; ok {
		return item.val, true
	}
	return nil, false
}

// Add -
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = CacheEntry{
		val:       value,
		createdAt: time.Now(),
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
