package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/analytics/entity"
	mock_analytics_uc "github.com/abmid/icanvas-analytics/internal/analytics/usecase/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestGetBestCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.TODO()
	mockAnalyticsUC := mock_analytics_uc.NewMockAnalyticsUseCase(ctrl)
	filter := entity.FilterAnalytics{
		AccountID: 1,
	}
	list := []entity.AnalyticsCourse{
		{ID: 1, CourseName: "Course Test"},
	}
	mockAnalyticsUC.EXPECT().FindBestCourseByFilter(ctx, filter).Return(list, nil)
	g := gin.Default()
	gr := g.Group("/")
	NewHandler("analytics", gr, mockAnalyticsUC)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/analytics/courses?account_id=1", nil)
	g.ServeHTTP(w, req)

	var result []entity.AnalyticsCourse
	json.NewDecoder(w.Body).Decode(&result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(result), len(list))
}
