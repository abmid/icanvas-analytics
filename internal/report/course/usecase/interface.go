/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

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
