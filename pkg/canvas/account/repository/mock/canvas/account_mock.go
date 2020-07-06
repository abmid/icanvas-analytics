package canvas

import (
	"net/http"

	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	echo "github.com/labstack/echo/v4"
)

func ListAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		var ListAccount []entity.Account
		ListAccount = []entity.Account{
			{ID: 1},
			{ID: 2},
		}
		return c.JSON(http.StatusOK, ListAccount)
	}
}
