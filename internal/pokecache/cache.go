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
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration, entries *map[string]cacheEntry) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	if entries != nil {
		cache.entries = *entries
	}

	go cache.reapLoop()

	return cache
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()

	for now := range ticker.C {
		cache.mu.Lock()

		for key, entry := range cache.entries {
			if now.Sub(entry.createdAt) > cache.interval {
				delete(cache.entries, key)
			}
		}

		cache.mu.Unlock()
	}

}

func (cache *Cache) Add(key string, val []byte) *Cache {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	existingCache, ok := cache.entries[key]

	if ok == false {
		cache.entries[key] = cacheEntry{
			val:       val,
			createdAt: time.Now(),
		}

		return cache
	}

	existingCache.val = val
	existingCache.createdAt = time.Now()

	return cache
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	existingCache, ok := cache.entries[key]

	if ok == false {
		return nil, false
	}

	return existingCache.val, true
}
