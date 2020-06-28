package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareAuthJWT(JWTKey string) echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(JWTKey),
	}
	midJWT := middleware.JWTWithConfig(config)

	return midJWT
}
