package user

import (
	"context"

	usercore "github.com/madkingxxx/backend-test/internal/core/user"
)

type RepoI interface {
	Get(ctx context.Context, ID int) (usercore.User, error)
	TopUp(ctx context.Context, ID int, amount float64) (usercore.User, error)
	Withdraw(ctx context.Context, ID int, amount float64) (usercore.User, error)
}

type Service struct {
	repo RepoI
}

func New(repo RepoI) *Service {
	return &Service{
		repo: repo,
	}
}

// Get retrieves a user from the repository.
func (service *Service) Get(ctx context.Context, ID int) (usercore.User, error) {
	return service.repo.Get(ctx, ID)
}

// TopUp increases a user's balance by the specified amount.
func (service *Service) TopUp(ctx context.Context, ID int, amount float64) (usercore.User, error) {
	user, err := service.repo.TopUp(ctx, ID, amount)
	if err != nil {
		return usercore.User{}, err
	}

	return user, nil
}

// Withdraw decreases a user's balance by the specified amount.
func (service *Service) Withdraw(ctx context.Context, ID int, amount float64) (usercore.User, error) {
	user, err := service.repo.Withdraw(ctx, ID, amount)
	if err != nil {
		return usercore.User{}, err
	}

	return user, nil
}
