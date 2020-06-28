package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/validation"
	"github.com/abmid/icanvas-analytics/pkg/analytics/entity"
	mock_analytics_uc "github.com/abmid/icanvas-analytics/pkg/analytics/usecase/mock"
	"github.com/abmid/icanvas-analytics/pkg/auth"
	echo "github.com/labstack/echo/v4"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestGetBestCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.TODO()
	// Mock Analytics
	mockAnalyticsUC := mock_analytics_uc.NewMockAnalyticsUseCase(ctrl)
	filter := entity.FilterAnalytics{
		AccountID: 1,
	}
	list := []entity.AnalyticsCourse{
		{ID: 1, CourseName: "Course Test"},
	}
	mockAnalyticsUC.EXPECT().FindBestCourseByFilter(ctx, filter).Return(list, nil)
	// Init Echo
	g := echo.New()
	validation.AlphaValidation(g)
	// Set Routing and Handler
	gr := g.Group("/v1")
	NewHandler("/analytics", gr, "super-secret", mockAnalyticsUC)
	// Generate Token
	token := auth.GenerateTokenDummy(1, "super-secret")
	w := httptest.NewRecorder()
	f := make(url.Values)
	f.Set("account_id", "1")
	req, _ := http.NewRequest("GET", "/v1/analytics/courses?"+f.Encode(), nil)
	req.Header.Add(echo.HeaderAuthorization, "Bearer "+token)
	g.ServeHTTP(w, req)

	var result []entity.AnalyticsCourse
	json.NewDecoder(w.Body).Decode(&result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(result), len(list))
}
