package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/analytics/entity"
)

type AnalyticsUseCase interface {
	FindBestCourseByFilter(ctx context.Context, filter entity.FilterAnalytics) ([]entity.AnalyticsCourse, error)
}
