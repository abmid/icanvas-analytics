package repository

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type EnrollmentRepositoryPG interface {
	Create(ctx context.Context, reportEnroll *entity.ReportEnrollment) error
	Read(ctx context.Context) ([]entity.ReportEnrollment, error)
	Update(ctx context.Context, reportEnroll *entity.ReportEnrollment) error
	Delete(ctx context.Context, reportEnrollID uint32) error
	FindFilter(ctx context.Context, filter entity.ReportEnrollment) ([]entity.ReportEnrollment, error)
	FindFirstByFilter(ctx context.Context, filter entity.ReportEnrollment) (entity.ReportEnrollment, error)
}