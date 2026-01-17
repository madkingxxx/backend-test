package skinport

import (
	"context"
	"slices"
	"testing"

	"github.com/madkingxxx/backend-test/internal/core/skinport"
)

func TestCacheAll(t *testing.T) {
	inmem := New()
	ctx := context.Background()

	items := []skinport.Item{
		{MarketHashName: "item1"},
		{MarketHashName: "item2"},
	}

	inmem.CacheAll(ctx, items)

	retrievedItems := inmem.GetAll(ctx)

	if len(retrievedItems) != len(items) {
		t.Fatalf("expected %d items, got %d", len(items), len(retrievedItems))
	}

	if !slices.Equal(retrievedItems, items) {
		t.Fatalf("expected items %v, got %v", items, retrievedItems)
	}
}
