package data

import (
	"errors"
	"sync"
	"time"
)

type fetchFunc[V any] func() ([]V, error)
type keyExtractFunc[K comparable, V any] func(V) K

var (
	ErrorNotFound = errors.New("record not found")
)

type BulkCache[K comparable, V any] struct {
	records      map[K]*V
	ttlSeconds   uint
	lastLoad     time.Time
	fetcher      fetchFunc[V]
	keyExtractor keyExtractFunc[K, V]
	mutex        *sync.RWMutex
}

func NewBulkCache[K comparable, V any](
	fetcher fetchFunc[V], keyExtractor keyExtractFunc[K, V], ttlSeconds uint) BulkCache[K, V] {
	return BulkCache[K, V]{
		records:      make(map[K]*V),
		ttlSeconds:   ttlSeconds,
		lastLoad:     time.Date(0, 0, 0, 0, 0, 0, 0, time.Local),
		fetcher:      fetcher,
		keyExtractor: keyExtractor,
		mutex:        &sync.RWMutex{},
	}
}

func (c *BulkCache[K, V]) IsFresh() bool {
	return c.lastLoad.Add(time.Duration(c.ttlSeconds) * time.Second).After(time.Now())
}

func (c *BulkCache[K, V]) Get(key K) (*V, error) {
	c.mutex.RLock()
	val, ok := c.records[key]
	c.mutex.RUnlock()
	if ok {
		return val, nil
	}
	err := c.Load()
	if err != nil {
		return nil, err
	}
	c.mutex.RLock()
	val, ok = c.records[key]
	c.mutex.RUnlock()
	if ok {
		return val, nil
	}
	return nil, ErrorNotFound
}

func (c *BulkCache[K, V]) All() []V {
	c.mutex.RLock()
	var records []V
	for _, v := range c.records {
		records = append(records, *v)
	}
	c.mutex.RUnlock()
	return records
}

func (c *BulkCache[K, V]) Load() error {
	records, err := c.fetcher()
	if err != nil {
		return err
	}
	c.Flush()
	for _, record := range records {
		key := c.keyExtractor(record)
		c.mutex.Lock()
		c.records[key] = &record
		c.mutex.Unlock()
	}
	c.lastLoad = time.Now()
	return nil
}

func (c *BulkCache[K, V]) Flush() {
	c.records = make(map[K]*V)
	c.lastLoad = time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
}
