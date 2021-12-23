package memcache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type MemStore struct {
	cache *cache.Cache
}

func NewMemStore(cache *cache.Cache) *MemStore {
	return &MemStore{
		cache: cache,
	}
}

func (store *MemStore) Get(key string) (interface{}, bool) {
	return store.cache.Get(key)
}
func (store *MemStore) Set(key string, value interface{}, d time.Duration) {
	store.cache.Set(key, value, d)
}
func (store *MemStore) Delete(key string) {
	store.cache.Delete(key)
}
