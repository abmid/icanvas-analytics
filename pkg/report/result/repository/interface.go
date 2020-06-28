/*
 * File Created: Thursday, 4th June 2020 3:00:53 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import (
	"context"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"
)

type ReportResultRepository interface {
	Update(ctx context.Context, result *entity.ReportResult) error
	Create(ctx context.Context, result *entity.ReportResult) error
	CreateOrUpdateByCourseReportID(ctx context.Context, reportResult *entity.ReportResult) error
}
