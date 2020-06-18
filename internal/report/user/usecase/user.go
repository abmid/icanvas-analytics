/*
 * File Created: Saturday, 6th June 2020 12:13:04 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

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
