package kv

import "time"

type KV interface {
	Set(key, value []byte) error
	SetWithTTL(key, value []byte, duration time.Duration) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	GCTicker(discardRatio float64, duration time.Duration)
	Close() error
}
