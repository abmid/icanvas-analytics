/*
 * File Created: Thursday, 4th June 2020 3:17:34 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"
	"github.com/abmid/icanvas-analytics/pkg/report/result/repository"
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
