package port

import (
	"github.com/labstack/echo/v4"
	errorscore "github.com/madkingxxx/backend-test/internal/core/errors"
	usercore "github.com/madkingxxx/backend-test/internal/core/user"
)

type GetUserRequest struct {
	ID int `json:"id" binding:"required" param:"id" validate:"required,numeric,min=1"`
}

type GetUserResponse struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

func ParseGetUserRequest(c echo.Context) (GetUserRequest, error) {
	var request GetUserRequest
	if err := c.Bind(&request); err != nil {
		return request, errorscore.ErrBinding
	}
	if err := c.Validate(&request); err != nil {
		return request, errorscore.ErrValidation
	}
	return request, nil
}

func Convert(user usercore.User) GetUserResponse {
	return GetUserResponse{
		ID:      user.ID,
		Balance: user.Balance,
	}
}
