package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type ReportResultUseCase interface {
	CreateOrUpdateByCourseReportID(ctx context.Context, reportResult *entity.ReportResult) error
}
