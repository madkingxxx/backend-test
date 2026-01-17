package user

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	usercore "github.com/madkingxxx/backend-test/internal/core/user"
	"github.com/madkingxxx/backend-test/internal/driver/http/user/port"
)

type usecaseI interface {
	Get(ctx context.Context, ID int) (usercore.User, error)
	TopUp(ctx context.Context, ID int, amount float64) (usercore.User, error)
	Purchase(ctx context.Context, ID int, marketHashName string) (usercore.User, error)
}

type Controller struct {
	usecase usecaseI
}

func New(usecase usecaseI) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

// GetUser godoc
// @Summary      Get User
// @Description  get user by ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  port.GetUserResponse
// @Failure      400  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /user/{id} [get]
func (ctr *Controller) GetUser(c echo.Context) error {
	request, err := port.ParseGetUserRequest(c)
	if err != nil {
		return err
	}

	response, err := ctr.usecase.Get(c.Request().Context(), request.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, port.Convert(response))
}

// TopUp godoc
// @Summary      Top Up
// @Description  top up user balance
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        amount   path      float64  true  "Amount"
// @Success      200  {object}  port.GetUserResponse
// @Failure      400  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /user/{id}/top-up/{amount} [get]
func (ctr *Controller) TopUp(c echo.Context) error {
	request, err := port.ParseTopUpRequest(c)
	if err != nil {
		return err
	}

	response, err := ctr.usecase.TopUp(c.Request().Context(), request.ID, request.Amount)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, port.Convert(response))
}

// Purchase godoc
// @Summary      Purchase
// @Description  purchase item
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        amount   path      float64  true  "Amount"
// @Success      200  {object}  port.GetUserResponse
// @Failure      400  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /user/{id}/purchase [POST]
func (ctr *Controller) Purchase(c echo.Context) error {
	request, err := port.ParsePurchaseRequest(c)
	if err != nil {
		return err
	}

	response, err := ctr.usecase.Purchase(c.Request().Context(), request.UserID, request.MarketHashName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, port.Convert(response))
}
