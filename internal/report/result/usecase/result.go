package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
	"github.com/abmid/icanvas-analytics/internal/report/result/repository"
)

type reportResultUseCase struct {
	RepoReportResult repository.ReportResultRepository
}

func NewReportResultUseCase(repoReportResult repository.ReportResultRepository) *reportResultUseCase {
	return &reportResultUseCase{
		RepoReportResult: repoReportResult,
	}
}

func (uc *reportResultUseCase) CreateOrUpdateByCourseReportID(ctx context.Context, reportResult *entity.ReportResult) error {
	err := uc.RepoReportResult.CreateOrUpdateByCourseReportID(ctx, reportResult)
	if err != nil {
		return err
	}
	return nil
}
