package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    *sync.Mutex
	store map[string]CacheItem
}

type CacheItem struct {
	value     any
	ttl       time.Duration
	createdAt time.Time
}

func New() *Cache {
	return &Cache{
		store: make(map[string]CacheItem),
		mu:    new(sync.Mutex),
	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = CacheItem{
		value:     value,
		ttl:       ttl,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.store[key]
	if !ok {
		return nil, false
	}

	if time.Now().After(v.createdAt.Add(v.ttl)) {
		delete(c.store, key)
		return nil, false
	}

	return v.value, ok
}

func (c Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.store, key)
}
