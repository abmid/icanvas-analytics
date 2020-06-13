package repository

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type AssigmentRepositoryPG interface {
	Create(ctx context.Context, assigment *entity.ReportAssigment) error
	Read(ctx context.Context) ([]entity.ReportAssigment, error)
	Update(ctx context.Context, assigment *entity.ReportAssigment) error
	Delete(ctx context.Context, reportAssigmentID uint32) error
	CreateOrUpdateByCourseReportID(ctx context.Context, assigment *entity.ReportAssigment) error
	FindFilter(ctx context.Context, filter entity.ReportAssigment) ([]entity.ReportAssigment, error)
	FindFirstByFilter(ctx context.Context, filter entity.ReportAssigment) (entity.ReportAssigment, error)
}