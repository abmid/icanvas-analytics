package repository

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type ReportResultRepository interface {
	Update(ctx context.Context, result *entity.ReportResult) error
	Create(ctx context.Context, result *entity.ReportResult) error
	CreateOrUpdateByCourseReportID(ctx context.Context, reportResult *entity.ReportResult) error
}
