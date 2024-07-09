package main

import (
	"sync"
	"time"
)

type Cache struct {
	store map[string]cacheEntry
	mu    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		store: map[string]cacheEntry{},
		mu:    &sync.RWMutex{},
	}

	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.reapLoop(interval)
		}
	}()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = cacheEntry{val: val, createdAt: time.Now().UTC()}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.store[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ago := time.Now().UTC().Add(-interval)
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.store {
		if entry.createdAt.Before(ago) {
			delete(c.store, key)
		}
	}
}
