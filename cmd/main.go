package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/madkingxxx/backend-test/internal/commands"
	"github.com/madkingxxx/backend-test/internal/config"
	"github.com/madkingxxx/backend-test/internal/connections/postgresql"
	"github.com/madkingxxx/backend-test/internal/driven"
	"github.com/madkingxxx/backend-test/internal/driver"
	"github.com/madkingxxx/backend-test/internal/migrator"
	"github.com/madkingxxx/backend-test/internal/server"
	"github.com/madkingxxx/backend-test/internal/service"
	"github.com/madkingxxx/backend-test/internal/usecase"
	"github.com/madkingxxx/backend-test/internal/utils"
	"go.uber.org/zap"
)

func main() {
	var err error
	time.Local, err = time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.New()
	utils.InitializeLogger(cfg.LogLevel)

	// Postgres connection
	pgPool := postgresql.New(ctx, cfg)
	defer pgPool.Close()

	// Migrator
	migratorService := migrator.New(utils.Logger.GetLogger(), cfg)
	if err := migratorService.Apply(); err != nil {
		utils.Logger.Fatal(ctx, "failed to apply migrations", zap.Error(err))
	}

	// Driven
	drivenManager := driven.New(cfg, pgPool)

	// Services
	serviceManager := service.New(drivenManager)

	// UseCases
	usecaseManager := usecase.New(serviceManager)

	// driver
	driverManager := driver.New(usecaseManager)

	// Server
	httpServer := server.New(cfg, driverManager)

	// Commands (Cron)
	scheduler := commands.NewScheduler()

	itemsCommand := commands.NewItemsCommand(cfg.ItemsCronExpression, serviceManager.Skinport)
	scheduler.Register(itemsCommand)

	scheduler.Start(ctx)

	// Start Server
	if err := httpServer.Run(ctx); err != nil {
		utils.Logger.Fatal(ctx, "server failed", zap.Error(err))
	}
}
