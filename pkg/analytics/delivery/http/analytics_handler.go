/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package http

import (
	"context"
	"net/http"

	"github.com/abmid/icanvas-analytics/pkg/analytics/entity"
	"github.com/abmid/icanvas-analytics/pkg/analytics/usecase"
	"github.com/abmid/icanvas-analytics/pkg/auth"
	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
)

type AnalyticsHandler struct {
	AUC usecase.AnalyticsUseCase
}

type ResponseError struct {
	Message string `json:"message"`
}

func (AH *AnalyticsHandler) GetBestCourse() echo.HandlerFunc {
	return func(c echo.Context) error {
		filter := new(entity.FilterAnalytics)
		/**
		* Bind query to Struct
		 */
		err := c.Bind(filter)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusBadRequest, ResponseError{Message: "Failed parameter"})
		}
		/**
		* Validate
		 */
		if err := c.Validate(filter); err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}
		ctx := context.TODO()
		res, err := AH.AUC.FindBestCourseByFilter(ctx, *filter)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusBadRequest, ResponseError{Message: "Failed to get resources"})

		}
		if res == nil {
			return c.JSON(http.StatusOK, ResponseError{Message: "Not found"})

		}
		return c.JSON(http.StatusOK, res)
	}
}

func NewHandler(path string, g *echo.Group, JWTKey string, articleUC usecase.AnalyticsUseCase) {
	handler := AnalyticsHandler{
		AUC: articleUC,
	}
	r := g.Group(path)
	r.Use(auth.MiddlewareAuthJWT(JWTKey))
	r.GET("/courses", handler.GetBestCourse())
}
