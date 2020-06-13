package repository

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type ReportResultRepository interface {
	Update(ctx context.Context, reportUser *entity.ReportUser) error
	Create(ctx context.Context, reportUser *entity.ReportUser) error
	CreateOrUpdateByCourseReportID(ctx context.Context, reportUser *entity.ReportUser) error
}
