package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/assigment/repository"
	"github.com/abmid/icanvas-analytics/internal/report/entity"
	report "github.com/abmid/icanvas-analytics/internal/report/entity"
)

type reportAssigmentUseCase struct {
	RepoReportAssgiment repository.AssigmentRepositoryPG
}

func NewReportAssigmentUseCase(repoReportAssigment repository.AssigmentRepositoryPG) *reportAssigmentUseCase {
	return &reportAssigmentUseCase{
		RepoReportAssgiment: repoReportAssigment,
	}
}

func (useCase *reportAssigmentUseCase) Create(ctx context.Context, reportAss *report.ReportAssigment) error {
	err := useCase.RepoReportAssgiment.Create(ctx, reportAss)
	if err != nil {
		return err
	}

	return nil
}

func (useCase *reportAssigmentUseCase) Read(ctx context.Context) ([]report.ReportAssigment, error) {
	results, err := useCase.RepoReportAssgiment.Read(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (useCase *reportAssigmentUseCase) Update(ctx context.Context, reportAss *report.ReportAssigment) error {
	err := useCase.RepoReportAssgiment.Update(ctx, reportAss)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *reportAssigmentUseCase) Delete(ctx context.Context, reportAssID uint32) error {
	err := useCase.RepoReportAssgiment.Delete(ctx, reportAssID)
	if err != nil {
		return err
	}
	return nil
}

// ! DEPRECATED
func (useCase *reportAssigmentUseCase) CreateOrUpdateByCourseReportID(ctx context.Context, assigment *entity.ReportAssigment) error {
	err := useCase.RepoReportAssgiment.CreateOrUpdateByCourseReportID(ctx, assigment)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *reportAssigmentUseCase) CreateOrUpdateByFilter(ctx context.Context, filter entity.ReportAssigment, assigment *entity.ReportAssigment) error {
	findReportAss, err := useCase.RepoReportAssgiment.FindFirstByFilter(ctx, filter)
	if err != nil {
		return err
	}
	assigment.ID = findReportAss.ID
	if findReportAss.ID == 0 {
		err := useCase.RepoReportAssgiment.Create(ctx, assigment)
		if err != nil {
			return err
		}
	} else {
		err := useCase.RepoReportAssgiment.Update(ctx, assigment)
		if err != nil {
			return err
		}
	}
	return nil
}

func (useCase *reportAssigmentUseCase) FindFilter(ctx context.Context, filter entity.ReportAssigment) (res []entity.ReportAssigment, err error) {

	find, err := useCase.RepoReportAssgiment.FindFilter(ctx, filter)
	if err != nil {
		return res, err
	}
	res = find
	return res, nil
}
