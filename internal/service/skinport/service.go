package skinport

import (
	"context"

	"github.com/madkingxxx/backend-test/internal/core/skinport"
	"github.com/madkingxxx/backend-test/internal/utils"
	"go.uber.org/zap"
)

type inMemI interface {
	Get(_ context.Context, hashName string) (skinport.Item, error)
	CacheAll(_ context.Context, item []skinport.Item)
	GetAll(_ context.Context) []skinport.Item
}

type senderI interface {
	GetAllItems(ctx context.Context) ([]skinport.Item, error)
}

type Service struct {
	inMemCache inMemI
	sender     senderI
}

func New(inMemCache inMemI, sender senderI) *Service {
	return &Service{
		inMemCache: inMemCache,
		sender:     sender,
	}
}

// Get retrieves an item from the in-memory cache.
func (s *Service) Get(ctx context.Context, hashName string) (skinport.Item, error) {
	return s.inMemCache.Get(ctx, hashName)
}

// GetAllItems retrieves all items from the in-memory cache.
func (s *Service) GetAllItems(ctx context.Context) []skinport.Item {
	return s.inMemCache.GetAll(ctx)
}

func (s *Service) getAllItemsHTTP(ctx context.Context) ([]skinport.Item, error) {
	items, err := s.sender.GetAllItems(ctx)
	if err != nil {
		utils.Logger.Error(ctx, "failed to get all items", zap.Error(err))
		return nil, err
	}
	return items, nil
}

func (s *Service) cacheAllItems(ctx context.Context, items []skinport.Item) {
	s.inMemCache.CacheAll(ctx, items)
}

// Cache fetches all items from the external HTTP sender and caches them in the in-memory cache.
func (s *Service) Cache(ctx context.Context) error {
	items, err := s.getAllItemsHTTP(ctx)
	if err != nil {
		return err
	}

	s.cacheAllItems(ctx, items)

	return nil
}
