package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/analytics/entity"
	mock_analytics_uc "github.com/abmid/icanvas-analytics/internal/analytics/usecase/mock"
	"github.com/abmid/icanvas-analytics/internal/validation"
	echo "github.com/labstack/echo/v4"

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
	g := echo.New()
	validation.AlphaValidation(g)
	gr := g.Group("/v1")
	NewHandler("/analytics", gr, mockAnalyticsUC)
	w := httptest.NewRecorder()
	f := make(url.Values)
	f.Set("account_id", "1")
	req, _ := http.NewRequest("GET", "/v1/analytics/courses?"+f.Encode(), nil)
	g.ServeHTTP(w, req)

	t.Log(w.Body.String())
	t.Fatalf("")

	var result []entity.AnalyticsCourse
	json.NewDecoder(w.Body).Decode(&result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(result), len(list))
}
