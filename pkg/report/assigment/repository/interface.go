/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import (
	"context"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"
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
