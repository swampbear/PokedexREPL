package pokecache

import (
	"sync"
	"time"
)

var mu sync.Mutex

type Cache struct {
	interval   time.Duration
	cacheEntry map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{interval, map[string]cacheEntry{}}
	go c.readLoop()
	return c
}

func (cache *Cache) Add(key string, val []byte) {
	mu.Lock()
	newCacheEntry := cacheEntry{time.Now(), val}
	cache.cacheEntry[key] = newCacheEntry
	mu.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, exits := cache.cacheEntry[key]
	if exits {
		return cacheEntry.val, true
	} else {
		return nil, false
	}
}

func (cache *Cache) readLoop() {
	ticker := time.NewTicker(cache.interval)
	for range ticker.C {
		for key, val := range cache.cacheEntry {
			if time.Since(val.createdAt) > cache.interval {
				delete(cache.cacheEntry, key)
			}
		}
	}
}
