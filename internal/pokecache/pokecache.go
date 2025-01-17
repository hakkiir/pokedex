package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.Mutex
	entry map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	c.entry[key] = cacheEntry{createdAt: time.Now(), val: val}
	defer c.mu.Unlock()
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if e, ok := c.entry[key]; !ok {
		return nil, false
	} else {
		return e.val, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()
		for key, val := range c.entry {
			if time.Since(val.createdAt) > interval {
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()
	}
}
