package http

import (
	"context"
	"net/http"

	"github.com/abmid/icanvas-analytics/internal/analytics/entity"
	"github.com/abmid/icanvas-analytics/internal/analytics/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AnalyticsHandler struct {
	AUC usecase.AnalyticsUseCase
}

type ResponseError struct {
	Message string `json:"message"`
}

func (AH *AnalyticsHandler) GetBestCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter entity.FilterAnalytics
		err := c.ShouldBind(&filter)
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, ResponseError{Message: "Parameter Salah"})
			return
		}
		ctx := context.TODO()
		res, err := AH.AUC.FindBestCourseByFilter(ctx, filter)
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, ResponseError{Message: "Data tidak ditemukan"})
			return
		}
		if res == nil {
			c.JSON(http.StatusOK, gin.H{
				"messages": "Not Found",
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func NewHandler(path string, g *gin.RouterGroup, articleUC usecase.AnalyticsUseCase) {
	handler := AnalyticsHandler{
		AUC: articleUC,
	}
	r := g.Group(path)
	r.GET("courses", handler.GetBestCourse())
}
