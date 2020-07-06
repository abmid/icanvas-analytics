package http

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type ResponseError struct {
	Message string `json:"message"`
}

func (h *CanvasAccountHandler) ListAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountID := c.Param("id")
		parseAccountID, err := strconv.Atoi(accountID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}
		res, err := h.AccountUseCase.ListAccount(uint32(parseAccountID))
		if err != nil {
			return c.JSON(http.StatusConflict, ResponseError{Message: err.Error()})
		}

		return c.JSON(http.StatusOK, res)
	}
}
