package skinport

import (
	"context"

	errorscore "github.com/madkingxxx/backend-test/internal/core/errors"
	"github.com/madkingxxx/backend-test/internal/core/skinport"
)

// Get retrieves an item from the in-memory cache.
func (inmem *Inmem) Get(_ context.Context, hashName string) (skinport.Item, error) {
	inmem.mu.RLock()
	defer inmem.mu.RUnlock()

	value, ok := inmem.cache[hashName]
	if !ok {
		return skinport.Item{}, errorscore.ErrNotFound
	}
	return value, nil
}

// CacheAll caches all items in the in-memory cache.
func (inmem *Inmem) CacheAll(_ context.Context, items []skinport.Item) {
	newCache := make(map[string]skinport.Item, len(items))
	for _, it := range items {
		newCache[it.MarketHashName] = it
	}

	inmem.mu.Lock()
	inmem.cache = newCache
	inmem.mu.Unlock()
}

// GetAll retrieves all items from the in-memory cache.
func (inmem *Inmem) GetAll(_ context.Context) []skinport.Item {
	inmem.mu.RLock()
	defer inmem.mu.RUnlock()

	var result []skinport.Item
	for _, item := range inmem.cache {
		result = append(result, item)
	}

	return result
}
