package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
	"github.com/abmid/icanvas-analytics/internal/report/user/repository"
)

type userUseCase struct {
	repoReportUser repository.ReportResultRepository
}

func NewReportUserUseCase(repoReportUser repository.ReportResultRepository) *userUseCase {
	return &userUseCase{
		repoReportUser: repoReportUser,
	}
}

func (uc *userUseCase) CreateOrUpdateByCourseReportID(ctx context.Context, reportUser *entity.ReportUser) error {
	err := uc.repoReportUser.CreateOrUpdateByCourseReportID(ctx, reportUser)
	if err != nil {
		return err
	}

	return nil

}
