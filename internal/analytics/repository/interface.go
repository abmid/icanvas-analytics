package repository

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/analytics/entity"
)

type AnalyticsRepository interface {
	FindBestCourseByFilter(ctx context.Context, filter entity.FilterAnalytics) ([]entity.AnalyticsCourse, error)
}
