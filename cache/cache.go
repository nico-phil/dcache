package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu sync.RWMutex
	db map[string]CacheItem
}

type CacheItem struct {
	Value      []byte
	ExpiryTime time.Time
}

func (c *CacheItem) IsExpired() bool {
	return time.Now().After(c.ExpiryTime)
}

func NewCache() *Cache {
	return &Cache{
		db: make(map[string]CacheItem, 0),
	}
}

func (c *Cache) Set(key string, value []byte, ttl time.Duration) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if ttl <= 0 {
		ttl = time.Hour * 24 // default TTL of 24 hours
	}

	c.db[key] = CacheItem{
		Value:      value,
		ExpiryTime: time.Now().Add(ttl),
	}
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.db[k]
	if !ok {
		return nil, false
	}
	if v.IsExpired() {
		delete(c.db, k)
		return nil, false
	}
	return v.Value, true
}

func (c *Cache) Delete(k string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.db[k]
	if !ok {
		return fmt.Errorf("Cache-Delete: key does not exist: %s", k)
	}
	delete(c.db, k)
	return nil
}

func (c *Cache) StartEvictionLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			c.evictExpiredItems()
		}
	}()

}

func (c *Cache) evictExpiredItems() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.db {
		if v.IsExpired() {
			delete(c.db, k)
		}
	}
}
