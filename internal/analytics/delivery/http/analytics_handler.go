package http

import (
	"context"
	"net/http"

	"github.com/abmid/icanvas-analytics/internal/analytics/entity"
	"github.com/abmid/icanvas-analytics/internal/analytics/usecase"
	"github.com/labstack/echo/v4"

	"github.com/gin-gonic/gin"
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
			return c.JSON(http.StatusBadRequest, ResponseError{Message: "Data tidak ditemukan"})

		}
		if res == nil {
			return c.JSON(http.StatusOK, gin.H{
				"messages": "Not Found",
			})

		}
		return c.JSON(http.StatusOK, res)
	}
}

func NewHandler(path string, g *echo.Group, articleUC usecase.AnalyticsUseCase) {
	handler := AnalyticsHandler{
		AUC: articleUC,
	}
	r := g.Group(path)
	r.GET("/courses", handler.GetBestCourse())
}
