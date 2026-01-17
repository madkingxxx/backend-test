package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	errorscore "github.com/madkingxxx/backend-test/internal/core/errors"
	usercore "github.com/madkingxxx/backend-test/internal/core/user"
)

const (
	getUserQuery             = `SELECT balance FROM users WHERE id=$1`
	withdrawUserBalanceQuery = `UPDATE users SET balance = balance - $1 WHERE id = $2 AND balance >= $1 RETURNING balance`
	topUPUserBalanceQuery    = `UPDATE users SET balance = balance + $1 WHERE id = $2 RETURNING balance`
)

func (r *Repo) Get(ctx context.Context, id int) (usercore.User, error) {
	var balance float64
	err := r.pool.QueryRow(ctx, getUserQuery, id).Scan(&balance)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return usercore.User{}, errorscore.ErrNotFound
		}
		return usercore.User{}, err
	}
	return usercore.User{ID: id, Balance: float64(balance)}, nil
}

func (r *Repo) Withdraw(ctx context.Context, id int, amount float64) (usercore.User, error) {
	var newBalance float64
	err := r.pool.QueryRow(ctx, withdrawUserBalanceQuery, amount, id).Scan(&newBalance)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return usercore.User{}, errorscore.ErrInsufficientFunds
		}
		return usercore.User{}, err
	}
	return usercore.User{ID: id, Balance: newBalance}, nil
}

func (r *Repo) TopUp(ctx context.Context, id int, amount float64) (usercore.User, error) {
	var newBalance float64
	err := r.pool.QueryRow(ctx, topUPUserBalanceQuery, amount, id).Scan(&newBalance)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return usercore.User{}, errorscore.ErrNotFound
		}
		return usercore.User{}, err
	}
	return usercore.User{ID: id, Balance: newBalance}, nil
}
