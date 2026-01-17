package skinport

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/madkingxxx/backend-test/internal/core/skinport"
	"github.com/madkingxxx/backend-test/internal/driver/http/skinport/port"
)

type usecaseI interface {
	GetAllItems(ctx context.Context) []skinport.Item
}

type Controller struct {
	usecase usecaseI
}

func New(usecase usecaseI) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

// GetAllItems godoc
// @Summary      Get Items
// @Description  get items
// @Tags         items
// @Accept       json
// @Produce      json
// @Success      200  {array}   port.ItemResponse
// @Failure      500  {object}  HTTPError
// @Router       /items [get]
func (ctr *Controller) GetAllItems(c echo.Context) error {
	ctx := c.Request().Context()

	items := ctr.usecase.GetAllItems(ctx)

	return c.JSON(http.StatusOK, port.Convert(items))
}
