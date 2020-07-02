/*
 * File Created: Thursday, 18th June 2020 5:28:49 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */
package http

import (
	"github.com/abmid/icanvas-analytics/pkg/auth/register/usecase"
	"github.com/labstack/echo/v4"
)

type RegisterHandler struct {
	registerUC usecase.RegisterUseCase
}

func NewHandler(basePath string, g *echo.Group, registerUC usecase.RegisterUseCase) {

	handler := RegisterHandler{
		registerUC: registerUC,
	}

	r := g.Group(basePath)
	r.POST("/register", handler.Register())
	r.GET("/register/check", handler.RegisterCheck())
}
