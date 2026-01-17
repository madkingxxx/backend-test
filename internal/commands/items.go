package commands

import (
	"context"

	"github.com/madkingxxx/backend-test/internal/utils"
	"go.uber.org/zap"
)

type itemsServiceI interface {
	Cache(ctx context.Context) error
}

func NewItemsCommand(cronExpression string, service itemsServiceI) *Command {
	return NewCommand(cronExpression, "items", func(ctx context.Context) {
		utils.Logger.Info(ctx, "starting items update command")
		if err := service.Cache(ctx); err != nil {
			utils.Logger.Error(ctx, "failed to update items", zap.Error(err))
			return
		}
		utils.Logger.Info(ctx, "successfully updated items")
	})
}
