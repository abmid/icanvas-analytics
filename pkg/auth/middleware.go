package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func MiddlewareAuthJWT(JWTKey string) echo.MiddlewareFunc {

	config := middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(JWTKey),
		// BeforeFunc to handle both request token from Header or Cookie
		// Token source from cookie will be convert to Header
		BeforeFunc: func(c echo.Context) {
			cookieToken, err := c.Cookie("icanvas_token")
			logrus.Info(cookieToken)
			headerToken := c.Request().Header.Get(echo.HeaderAuthorization)
			logrus.Info(headerToken)
			if err == nil && cookieToken.Value != "" && headerToken == "" {
				c.Request().Header.Add(echo.HeaderAuthorization, "Bearer "+cookieToken.Value)
			}
		},
	}
	midJWT := middleware.JWTWithConfig(config)

	return midJWT
}
