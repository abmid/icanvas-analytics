package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type ReportCourseUseCase interface {
	Create(ctx context.Context, reportCourse *entity.ReportCourse) error
	Read(ctx context.Context) ([]entity.ReportCourse, error)
	FindFilter(ctx context.Context, filter entity.ReportCourse) ([]entity.ReportCourse, error)
	Update(ctx context.Context, reportCourse *entity.ReportCourse) error
	Delete(ctx context.Context, reportCourseID uint32) error
	CreateOrUpdateCourseID(ctx context.Context, reportCourse *entity.ReportCourse) error
}
