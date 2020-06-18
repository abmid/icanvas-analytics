/*
 * File Created: Thursday, 4th June 2020 3:16:20 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type ReportResultUseCase interface {
	CreateOrUpdateByCourseReportID(ctx context.Context, reportResult *entity.ReportResult) error
}
