package port

import (
	"github.com/labstack/echo/v4"
	errorscore "github.com/madkingxxx/backend-test/internal/core/errors"
)

type PurchaseRequest struct {
	UserID         int    `json:"user_id" binding:"required" param:"user_id" validate:"required,numeric,min=1"`
	MarketHashName string `json:"market_hash_name" binding:"required" param:"market_hash_name" validate:"required"`
}

func ParsePurchaseRequest(c echo.Context) (PurchaseRequest, error) {
	var request PurchaseRequest
	if err := c.Bind(&request); err != nil {
		return request, errorscore.ErrBinding
	}
	if err := c.Validate(&request); err != nil {
		return request, errorscore.ErrValidation
	}
	return request, nil
}
