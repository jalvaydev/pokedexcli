package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCacheEntry(val []byte) cacheEntry {
	return cacheEntry{createdAt: time.Now(), val: val}
}

type Cache struct {
	cache map[string]cacheEntry
	mutex sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
	}

	go cache.reapLoop(time.Duration(5) * time.Second)

	return &cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		if ticker.C != nil {
			c.mutex.Lock()
			for key, v := range c.cache {
				if interval <= time.Since(v.createdAt) {
					delete(c.cache, key)
				}
			}
			c.mutex.Unlock()
		}
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	entry := NewCacheEntry(val)
	c.cache[key] = entry
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	v, ok := c.cache[key]
	c.mutex.Unlock()
	return v.val, ok
}
