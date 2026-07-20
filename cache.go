package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu sync.Mutex
	db map[string][]byte
}

func NewCache() *Cache {
	return &Cache{
		db: make(map[string][]byte, 0),
	}
}

func (c *Cache) Add(key string, value []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.db[key]
	if !ok {
		c.db[key] = value
	}

	return fmt.Errorf("Cache-Add: key already exist")
}

func (c *Cache) Get(k string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.db[k]
	if !ok {
		return nil, fmt.Errorf("Cache-Get: key does not exist: %s", k)
	}

	return v, nil

}

func main() {

}
