/*
 * File Created: Wednesday, 1st July 2020 2:28:58 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package http

import (
	"net/http"

	"github.com/abmid/icanvas-analytics/internal/logger"
	"github.com/abmid/icanvas-analytics/pkg/auth"
	echo "github.com/labstack/echo/v4"
)

type LogoutHandler struct {
	Log *logger.LoggerWrap
}

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Status string `json:"status"`
}

func (LH *LogoutHandler) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := auth.WriteTokenCookie(c, "logout")
		if err != nil {
			LH.Log.Error(err)
			return c.JSON(http.StatusConflict, ResponseError{Message: "Failed logout !"})
		}

		return c.JSON(http.StatusOK, ResponseSuccess{Status: "OK"})
	}
}

func NewHandler(path string, g *echo.Group) {

	logger := logger.New()

	handler := LogoutHandler{
		Log: logger,
	}

	r := g.Group(path)
	r.POST("/logout", handler.Logout())
}
