/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"
)

type ReportEnrollmentUseCase interface {
	Create(ctx context.Context, reportEnroll *entity.ReportEnrollment) error
	Read(ctx context.Context) ([]entity.ReportEnrollment, error)
	Update(ctx context.Context, reportEnroll *entity.ReportEnrollment) error
	Delete(ctx context.Context, reportEnrollID uint32) error
	CreateOrUpdateByFilter(ctx context.Context, filter entity.ReportEnrollment, assigment *entity.ReportEnrollment) error
	FindFilter(ctx context.Context, filter entity.ReportEnrollment) ([]entity.ReportEnrollment, error)
}
