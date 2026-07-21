package cache

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu sync.RWMutex
	db map[string][]byte
}

func NewCache() *Cache {
	return &Cache{
		db: make(map[string][]byte, 0),
	}
}

func (c *Cache) Set(key string, value []byte) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	c.db[key] = value

}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.db[k]
	if !ok {
		return nil, false
	}
	return v, true
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
