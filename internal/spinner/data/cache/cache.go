package cache

import "time"

type CacheEntry struct {
	Data      []byte
	UpdatedAt time.Time
	ExpiresAt time.Time
}

type Cache interface {
	Get(key string) (*CacheEntry, bool)
	Put(key string, data []byte, expiresAt time.Time) error
	Delete(key string) error
}
