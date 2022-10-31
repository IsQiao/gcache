package gcache

import (
	"sync"
	"time"
)

const (
	default_CleanupInterval = time.Second
)

type GCache[T interface{}] struct {
	mu              sync.RWMutex
	cleanupInterval time.Duration
	expiration      time.Duration
	cacheMap        map[string]GCacheItem[T]
}

func New[T interface{}](cleanupInterval time.Duration, expiration time.Duration) *GCache[T] {
	gc := &GCache[T]{
		cleanupInterval: cleanupInterval,
		expiration:      expiration,
		cacheMap:        map[string]GCacheItem[T]{},
	}

	go gc.startClearJob()
	return gc
}

func NewDefault[T interface{}](expiration time.Duration) *GCache[T] {
	return New[T](default_CleanupInterval, expiration)
}

func (gc *GCache[T]) Set(key string, item T) {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	now := time.Now()
	gc.cacheMap[key] = GCacheItem[T]{
		expiredTime: now.Add(gc.expiration),
		item:        item,
	}
}

// Clear all cache
func (gc *GCache[T]) Flush() {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	gc.cacheMap = map[string]GCacheItem[T]{}
}

// Get one cache
func (gc *GCache[T]) Get(key string) *T {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	if val, ok := gc.cacheMap[key]; ok {
		return &val.item
	}

	return nil
}

// Delete one cache if existed
func (gc *GCache[T]) Delete(key string) {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	delete(gc.cacheMap, key)
}

func (gc *GCache[T]) startClearJob() {
	t := time.NewTicker(gc.cleanupInterval)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			gc.DeleteAllExpired()
		}
	}
}

// Delete all expired items.
func (gc *GCache[T]) DeleteAllExpired() {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	for cacheKey, cacheItem := range gc.cacheMap {
		if cacheItem.Expired() {
			delete(gc.cacheMap, cacheKey)
		}
	}
}
