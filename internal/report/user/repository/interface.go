/*
 * File Created: Saturday, 6th June 2020 10:48:28 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type ReportResultRepository interface {
	Update(ctx context.Context, reportUser *entity.ReportUser) error
	Create(ctx context.Context, reportUser *entity.ReportUser) error
	CreateOrUpdateByCourseReportID(ctx context.Context, reportUser *entity.ReportUser) error
}
