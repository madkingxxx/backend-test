package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/madkingxxx/backend-test/internal/config"
	"github.com/madkingxxx/backend-test/internal/driver"
	"github.com/madkingxxx/backend-test/internal/utils"
	"go.uber.org/zap"
)

type HTTPServer struct {
	echo             *echo.Echo
	cfg              *config.Config
	transportManager *driver.Config
}

func New(
	cfg *config.Config,
	transportManager *driver.Config,
) *HTTPServer {
	return &HTTPServer{
		echo:             echo.New(),
		cfg:              cfg,
		transportManager: transportManager,
	}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	s.echo.Validator = NewValidator()
	s.SetUpRoutes()
	s.echo.HTTPErrorHandler = customHTTPErrorHandler

	go func() {
		address := fmt.Sprintf(":%d", s.cfg.ServerPort)
		utils.Logger.Info(ctx, "starting server", zap.String("address", address))
		if err := s.echo.Start(address); err != nil && err != http.ErrServerClosed {
			utils.Logger.Fatal(ctx, "shutting down the server", zap.Error(err))
		}
	}()

	<-ctx.Done()

	utils.Logger.Info(ctx, "shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.echo.Shutdown(shutdownCtx); err != nil {
		return err
	}

	return nil
}
