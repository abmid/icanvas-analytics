/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/enrollment/repository"
	"github.com/abmid/icanvas-analytics/internal/report/entity"
	report "github.com/abmid/icanvas-analytics/internal/report/entity"
)

type reportEnrollUseCase struct {
	RepoReportEnroll repository.EnrollmentRepositoryPG
}

func NewReportEnrollUseCase(repoReportEnroll repository.EnrollmentRepositoryPG) *reportEnrollUseCase {
	return &reportEnrollUseCase{
		RepoReportEnroll: repoReportEnroll,
	}
}

func (useCase *reportEnrollUseCase) Create(ctx context.Context, reportEnroll *report.ReportEnrollment) error {
	err := useCase.RepoReportEnroll.Create(ctx, reportEnroll)
	if err != nil {
		return err
	}

	return nil
}

func (useCase *reportEnrollUseCase) Read(ctx context.Context) ([]report.ReportEnrollment, error) {
	results, err := useCase.RepoReportEnroll.Read(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (useCase *reportEnrollUseCase) Update(ctx context.Context, reportEnroll *report.ReportEnrollment) error {
	err := useCase.RepoReportEnroll.Update(ctx, reportEnroll)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *reportEnrollUseCase) Delete(ctx context.Context, reportEnrollID uint32) error {
	err := useCase.RepoReportEnroll.Delete(ctx, reportEnrollID)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *reportEnrollUseCase) CreateOrUpdateByFilter(ctx context.Context, filter report.ReportEnrollment, reportEnroll *report.ReportEnrollment) error {
	findReportEnroll, err := useCase.RepoReportEnroll.FindFirstByFilter(ctx, filter)
	if err != nil {
		return err
	}
	reportEnroll.ID = findReportEnroll.ID
	if findReportEnroll.ID == 0 {
		err := useCase.RepoReportEnroll.Create(ctx, reportEnroll)
		if err != nil {
			return err
		}
	} else {
		err := useCase.RepoReportEnroll.Update(ctx, reportEnroll)
		if err != nil {
			return err
		}
	}
	return nil
}

func (useCase *reportEnrollUseCase) FindFilter(ctx context.Context, filter entity.ReportEnrollment) ([]entity.ReportEnrollment, error) {
	results, err := useCase.RepoReportEnroll.FindFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return results, nil
}
