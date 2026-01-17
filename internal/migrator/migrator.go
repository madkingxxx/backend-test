package migrator

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/madkingxxx/backend-test/internal/config"
	"go.uber.org/zap"
)

type Migrator struct {
	logger *zap.Logger
	cfg    *config.Config
}

func New(logger *zap.Logger, cfg *config.Config) *Migrator {
	return &Migrator{
		logger: logger,
		cfg:    cfg,
	}
}

func (m *Migrator) Apply() error {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		m.cfg.DBUser,
		m.cfg.DBPassword,
		m.cfg.DBHost,
		m.cfg.DBPort,
		m.cfg.DBName,
		m.cfg.DBSSLMode,
	)

	var mig *migrate.Migrate
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		mig, err = migrate.New("file://migrations", dbURL)
		if err == nil {
			break
		}
		m.logger.Warn("Failed to connect to database for migration, retrying...", zap.Error(err), zap.Int("attempt", i+1))
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := mig.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			m.logger.Info("No database schema changes were applied")
			return nil
		}
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	m.logger.Info("Database migrations applied successfully")
	return nil
}
