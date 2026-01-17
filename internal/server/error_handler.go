package server

import (
	"context"
	"net/http"

	"errors"

	"github.com/labstack/echo/v4"

	errorscore "github.com/madkingxxx/backend-test/internal/core/errors"
	"github.com/madkingxxx/backend-test/internal/utils"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	var e *echo.HTTPError

	if !errors.As(err, &e) {
		switch {
		case errors.Is(err, context.Canceled):
			e = echo.NewHTTPError(499, "Client closed request")
		case errors.Is(err, errorscore.ErrNotFound):
			e = echo.NewHTTPError(http.StatusNotFound, err.Error())
		case errors.Is(err, errorscore.ErrInsufficientFunds):
			e = echo.NewHTTPError(http.StatusPaymentRequired, err.Error())
		case errors.Is(err, errorscore.ErrValidation):
			e = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		case errors.Is(err, errorscore.ErrBinding):
			e = echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		default:
			e = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	err = c.JSON(e.Code, map[string]any{
		"code":    e.Code,
		"message": e.Message,
	})
	if err != nil {
		utils.Logger.Error(c.Request().Context(), errors.New("customHTTPErrorHandler"))
	}
}
