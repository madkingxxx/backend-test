package skinport

import (
	"context"

	"github.com/madkingxxx/backend-test/internal/core/skinport"
)

type skinportServiceI interface {
	Get(ctx context.Context, hashName string) (skinport.Item, error)
	GetAllItems(ctx context.Context) []skinport.Item
}

type UseCase struct {
	skinportService skinportServiceI
}

func New(skinportService skinportServiceI) *UseCase {
	return &UseCase{
		skinportService: skinportService,
	}
}

// GetAllItems retrieves all items from the skinport service.
func (usecase *UseCase) GetAllItems(ctx context.Context) []skinport.Item {
	return usecase.skinportService.GetAllItems(ctx)
}

// Get retrieves an item from the skinport service.
func (usecase *UseCase) Get(ctx context.Context, hashName string) (skinport.Item, error) {
	return usecase.skinportService.Get(ctx, hashName)
}
