/*
 * File Created: Saturday, 6th June 2020 12:12:13 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"
)

type ReportUserUseCase interface {
	CreateOrUpdateByCourseReportID(ctx context.Context, reportUser *entity.ReportUser) error
}
