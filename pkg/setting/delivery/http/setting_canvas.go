/*
 * File Created: Tuesday, 28th July 2020 5:51:44 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package http

import (
	"net/http"

	"github.com/abmid/icanvas-analytics/internal/inerr"
	"github.com/abmid/icanvas-analytics/pkg/setting/entity"
	echo "github.com/labstack/echo/v4"
)

type FormDataCanvas struct {
	CanvasUrl   string `json:"canvas_url" form:"canvas_url" validate:"required"`
	CanvasToken string `json:"canvas_token" form:"canvas_token" validate:"required"`
}

func (SH *SettingHandler) CreateOrUpdateCanvas() echo.HandlerFunc {
	return func(c echo.Context) error {
		formData := new(FormDataCanvas)
		// Bind First
		if err := c.Bind(formData); err != nil {
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}
		// Validate
		if err := c.Validate(formData); err != nil {
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}

		// Save
		settings := []*entity.Setting{
			{Category: "canvas", Name: "url", Value: formData.CanvasUrl},
			{Category: "canvas", Name: "token", Value: formData.CanvasToken},
		}
		err := SH.SettingUseCase.CreateAll(settings)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
		}

		return c.JSON(http.StatusOK, ResponseSuccess{Message: "success"})
	}
}

func (SH *SettingHandler) ExistsCanvasConfig() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get From UseCase
		exists, url, token, err := SH.SettingUseCase.ExistsCanvasConfig()
		if err != nil {
			return c.JSON(http.StatusConflict, ResponseError{Message: err.Error()})
		}

		if !exists {
			return c.JSON(http.StatusConflict, ResponseError{Message: inerr.ErrNoCanvasConfig.Error()})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"url":   url,
			"token": token,
		})
	}
}
