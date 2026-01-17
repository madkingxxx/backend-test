package port

import (
	"github.com/labstack/echo/v4"
	errorscore "github.com/madkingxxx/backend-test/internal/core/errors"
)

func ParseTopUpRequest(c echo.Context) (TopUpRequest, error) {
	var request TopUpRequest
	if err := c.Bind(&request); err != nil {
		return request, errorscore.ErrBinding
	}
	if err := c.Validate(&request); err != nil {
		return request, errorscore.ErrValidation
	}
	return request, nil
}

type TopUpRequest struct {
	ID     int     `json:"id" binding:"required" param:"id" validate:"required,numeric,min=1"`
	Amount float64 `json:"amount" binding:"required" param:"amount" validate:"required,numeric,min=0.5"`
}
