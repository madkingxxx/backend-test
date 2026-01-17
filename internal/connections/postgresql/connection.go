package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/madkingxxx/backend-test/internal/config"
	"github.com/madkingxxx/backend-test/internal/utils"
	"go.uber.org/zap"
)

func New(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		utils.Logger.Fatal(ctx, "failed to parse postgres config", zap.Error(err))
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		utils.Logger.Fatal(ctx, "failed to connect to postgres", zap.Error(err))
	}

	if err := pool.Ping(ctx); err != nil {
		utils.Logger.Fatal(ctx, "failed to ping postgres", zap.Error(err))
	}

	utils.Logger.Info(ctx, "successfully connected to postgres")

	return pool
}
