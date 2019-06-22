package cache

import (
	"sync"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

type GoCache struct {
	gocache *gocache.Cache
	mutex   sync.Mutex
}

func OpenGoCache(defaultExpiration, cleanupInterval time.Duration) Cache {
	return &GoCache{
		gocache: gocache.New(defaultExpiration, cleanupInterval),
	}
}

func (db *GoCache) Get(key string) (interface{}, bool) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	return db.gocache.Get(key)
}

func (db *GoCache) Set(key string, value interface{}, duration time.Duration) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.gocache.Set(key, value, duration)
}

func (db *GoCache) Delete(key string) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.gocache.Set(key, "", time.Nanosecond)
}
