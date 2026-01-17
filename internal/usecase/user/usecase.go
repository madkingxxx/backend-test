package user

import (
	"context"

	"github.com/madkingxxx/backend-test/internal/core/skinport"
	usercore "github.com/madkingxxx/backend-test/internal/core/user"
)

type userServiceI interface {
	Get(ctx context.Context, id int) (usercore.User, error)
	Withdraw(ctx context.Context, id int, amount float64) (usercore.User, error)
	TopUp(ctx context.Context, id int, amount float64) (usercore.User, error)
}

type skinportServiceI interface {
	Get(ctx context.Context, hashName string) (skinport.Item, error)
}

type UseCase struct {
	userService     userServiceI
	skinportService skinportServiceI
}

func New(userService userServiceI, skinportService skinportServiceI) *UseCase {
	return &UseCase{
		userService:     userService,
		skinportService: skinportService,
	}
}

func (usecase *UseCase) Get(ctx context.Context, id int) (usercore.User, error) {
	return usecase.userService.Get(ctx, id)
}

func (usecase *UseCase) TopUp(ctx context.Context, id int, amount float64) (usercore.User, error) {
	return usecase.userService.TopUp(ctx, id, amount)
}

func (usecase *UseCase) Purchase(ctx context.Context, userID int, hashName string) (usercore.User, error) {
	item, err := usecase.skinportService.Get(ctx, hashName)
	if err != nil {
		return usercore.User{}, err
	}

	user, err := usecase.userService.Withdraw(ctx, userID, item.MinPrice)
	if err != nil {
		return usercore.User{}, err
	}

	return user, nil
}
