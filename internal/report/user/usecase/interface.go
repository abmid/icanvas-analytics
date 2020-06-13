package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type ReportUserUseCase interface {
	CreateOrUpdateByCourseReportID(ctx context.Context, reportUser *entity.ReportUser) error
}
