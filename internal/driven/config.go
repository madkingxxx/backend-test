package driven

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/madkingxxx/backend-test/internal/config"
	exthttp "github.com/madkingxxx/backend-test/internal/driven/ext_http"
	"github.com/madkingxxx/backend-test/internal/driven/inmem"
	"github.com/madkingxxx/backend-test/internal/driven/postgresql"
)

type Config struct {
	Inmem      *inmem.Config
	ExtHTTP    *exthttp.Config
	PostgreSQL *postgresql.Config
}

func New(cfg *config.Config, pgconn *pgxpool.Pool) *Config {
	return &Config{
		Inmem:      inmem.New(),
		ExtHTTP:    exthttp.New(cfg.SkinportAPIBaseURL),
		PostgreSQL: postgresql.New(pgconn),
	}
}
