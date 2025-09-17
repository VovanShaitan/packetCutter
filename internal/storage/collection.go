package storage

import (
	"sync"
)

type ResultCollection struct {
	mu   sync.RWMutex
	data map[string][]string
}

func NewResultCollection() *ResultCollection {
	return &ResultCollection{data: make(map[string][]string)}
}

func (rc *ResultCollection) Add(hexResult, target string) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.data[hexResult] = append(rc.data[hexResult], target)
}

func (rc *ResultCollection) Get(hexResult string) ([]string, bool) {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	targets, exists := rc.data[hexResult]
	return targets, exists
}

func (rc *ResultCollection) GetAll() map[string][]string {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	copy := make(map[string][]string, len(rc.data))
	for k, v := range rc.data {
		copy[k] = append([]string{}, v...)
	}
	return copy
}

func (rc *ResultCollection) Len() int {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return len(rc.data)
}

func (rc *ResultCollection) Count() int {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	total := 0
	for _, targets := range rc.data {
		total += len(targets)
	}
	return total
}
