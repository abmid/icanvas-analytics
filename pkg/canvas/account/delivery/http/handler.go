package http

import (
	"github.com/abmid/icanvas-analytics/pkg/auth"
	"github.com/abmid/icanvas-analytics/pkg/canvas/account/usecase"
	echo "github.com/labstack/echo/v4"
)

type CanvasAccountHandler struct {
	AccountUseCase usecase.AccountUseCase
}

func NewHandler(path string, g *echo.Group, JWTKey string, accountUseCase usecase.AccountUseCase) {
	handler := CanvasAccountHandler{
		AccountUseCase: accountUseCase,
	}
	r := g.Group(path)
	r.Use(auth.MiddlewareAuthJWT(JWTKey))
	r.GET("/accounts/:id/sub_accounts", handler.ListAccount())
}
