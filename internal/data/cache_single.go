package data

import (
	"sync"
	"time"
)

type fetchSingleFunc[K comparable, V any] func(K) (*V, error)

type CacheRecord[V any] struct {
	record   *V
	loadTime time.Time
}

type SingleCache[K comparable, V any] struct {
	records      map[K]CacheRecord[V]
	ttlSeconds   time.Duration
	fetcher      fetchSingleFunc[K, V]
	keyExtractor keyExtractFunc[K, V]
	mutex        *sync.RWMutex
}

func NewSingleCache[K comparable, V any](
	fetcher fetchSingleFunc[K, V], keyExtractor keyExtractFunc[K, V], ttlSeconds time.Duration) SingleCache[K, V] {
	return SingleCache[K, V]{
		records:      make(map[K]CacheRecord[V]),
		ttlSeconds:   ttlSeconds,
		fetcher:      fetcher,
		keyExtractor: keyExtractor,
		mutex:        &sync.RWMutex{},
	}
}

func (c *SingleCache[K, V]) IsFresh(record CacheRecord[V]) bool {
	return record.loadTime.Add(c.ttlSeconds).After(time.Now())
}

func (c *SingleCache[K, V]) Get(key K) (*V, error) {
	c.mutex.RLock()
	val, ok := c.records[key]
	c.mutex.RUnlock()
	if ok && c.IsFresh(val) {
		return val.record, nil
	}
	record, err := c.Load(key)
	if err != nil {
		return nil, err
	}
	if record != nil {
		return record, nil
	}

	return nil, ErrorNotFound
}

func (c *SingleCache[K, V]) Add(record V) {
	key := c.keyExtractor(record)
	c.mutex.Lock()
	c.records[key] = CacheRecord[V]{
		record:   &record,
		loadTime: time.Now(),
	}
	c.mutex.Unlock()
}

func (c *SingleCache[K, V]) Load(key K) (*V, error) {
	record, err := c.fetcher(key)
	if err != nil {
		return nil, err
	}
	if record != nil {
		c.Add(*record)
	}
	return record, err
}

func (c *SingleCache[K, V]) Flush() {
	c.records = make(map[K]CacheRecord[V])
}
