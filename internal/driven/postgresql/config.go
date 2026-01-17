package postgresql

import (
	"github.com/jackc/pgx/v5/pgxpool"
	userpostgres "github.com/madkingxxx/backend-test/internal/driven/postgresql/user/adapter"
)

type Config struct {
	User *userpostgres.Repo
}

func New(
	pgConn *pgxpool.Pool,
) *Config {
	return &Config{
		User: userpostgres.New(pgConn),
	}
}
